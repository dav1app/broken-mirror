package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

// separator is a private use character (U+E000) used as a separator
var separator = []byte{0xEE, 0x80, 0x80}

func main() {
	log.Printf("%+v", Main(map[string]interface{}{
		"username": "test",
		"password": "12345678",
	}))
}

// Main is the entry point for the serverless function
func Main(args map[string]interface{}) map[string]interface{} {
	response := make(map[string]interface{})

	username, okUsername := args["username"].(string)
	password, okPassword := args["password"].(string)

	if !okUsername || !okPassword {
		response["error"] = "username or password not provided or invalid"
		return createHTTPResponse(response, 400)
	}

	if len(password) < 8 || utf8.RuneCountInString(password) < 8 {
		response["error"] = "password must be at least 8 characters long"
		return createHTTPResponse(response, 400)
	}

	if len(password) > 64 || utf8.RuneCountInString(password) > 64 {
		response["error"] = "password must be at most 64 characters long"
		return createHTTPResponse(response, 400)
	}

	usernameBuffer := []byte(username)
	passwordBuffer := []byte(password)

	if len(passwordBuffer) > 72 {
		response["error"] = "password must be at most 72 bytes long"
		return createHTTPResponse(response, 400)
	}

	if bytes.Contains(usernameBuffer, separator) || bytes.Contains(passwordBuffer, separator) {
		response["error"] = "username or password contains invalid characters"
		return createHTTPResponse(response, 400)
	}

	hasher := sha1.New()
	hasher.Write(passwordBuffer)
	hashedSHA1Password := strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
	hashPrefix := hashedSHA1Password[:5] // First 5 chars
	hashSuffix := hashedSHA1Password[5:] // Rest of the characters

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + hashPrefix)
	if err != nil {
		response["error"] = "unable to make request to pwnedpasswords"
		return createHTTPResponse(response, 500)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != http.StatusOK {
		response["error"] = "unable to read response from pwnedpasswords"
		return createHTTPResponse(response, 500)
	}

	if strings.Contains(string(body), hashSuffix) {
		response["error"] = "password is leaked"
		return createHTTPResponse(response, 400)
	}

	combined := make([]byte, 0, len(usernameBuffer)+len(separator)+len(passwordBuffer))
	combined = append(combined, usernameBuffer...)
	combined = append(combined, separator...)
	combined = append(combined, passwordBuffer...)

	hasher256 := sha256.New()
	hasher256.Write(combined)
	hashedValue := hex.EncodeToString(hasher256.Sum(nil))

	response["hash"] = hashedValue

	return createHTTPResponse(response, 200)
}

// createHTTPResponse creates and returns a HTTP response
func createHTTPResponse(response map[string]interface{}, statusCode int) map[string]interface{} {
	jsonBody, err := json.Marshal(response)
	if err != nil {
		// Handle JSON marshalling error
		return map[string]interface{}{
			"headers":    map[string]interface{}{"Content-Type": "application/json"},
			"statusCode": 400,
			"body":       "{\"error\":\"Internal server error\"}",
		}
	}

	return map[string]interface{}{
		"headers":    map[string]interface{}{"Content-Type": "application/json"},
		"statusCode": statusCode,
		"body":       string(jsonBody),
	}
}
