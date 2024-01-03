# Bug Bounty

Find any bugs on this endpoint: 

```
curl -X POST https://faas-nyc1-2ef2e6cc.doserverless.co/api/v1/web/fn-ba8e141c-152f-46c2-9f1c-41dc67d7d55f/brokenmirror/brokenmirror \
-H "Content-Type: application/json" \
-d '{
  "username": "my-username",
  "password": "123456789!!%%ArAGoRn"
}'

{
  "hash": "4ea78f2aa045b6c5a9d858705f15db54293f4ac54c2264173eaf92a2fd1936ee"
}

```

Expected response in case of error:

```
curl -X POST https://faas-nyc1-2ef2e6cc.doserverless.co/api/v1/web/fn-ba8e141c-152f-46c2-9f1c-41dc67d7d55f/brokenmirror/brokenmirror \
-H "Content-Type: application/json" \
-d '{
  "username": "my-username",
  "password": "123456789"
}'

{
    "error": "password is leaked"
}
```

Here are the error messages in the provided code:

1. **"username or password not provided or invalid"** - This error message is returned when the `username` or `password` is not provided or is invalid.

2. **"password must be at least 8 characters long"** - This error message is returned when the length of the `password` is less than 8 characters.

3. **"password must be at most 64 characters long"** - This error message is returned when the length of the `password` is greater than 64 characters.

4. **"password must be at most 72 bytes long"** - This error message is returned when the length of the `password` in bytes is greater than 72 bytes.

5. **"username or password contains invalid characters"** - This error message is returned when the `username` or `password` contains invalid characters.

6. **"unable to make request to pwnedpasswords"** - This error message is returned when there is an error making a request to the "https://api.pwnedpasswords.com/range/" endpoint.

7. **"unable to read response from pwnedpasswords"** - This error message is returned when there is an error reading the response from the "https://api.pwnedpasswords.com/range/" endpoint.

8. **"password is leaked"** - This error message is returned when the `password` is found in the list of leaked passwords.


Any other error message can be considered a bug.
