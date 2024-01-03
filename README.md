# Bug Bounty

## Overview
Our service is designed to strengthen password security by preventing the usage of leaked and weak passwords. It ensures that users choose passwords that are not only robust but also have not been compromised in data breaches.

## Endpoint Description
This service provides an endpoint to validate passwords against a database of known leaked passwords and enforce strength requirements.

## Endpoint Usage
To use the service, make a POST request with a JSON payload containing the username and password:

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

1. **"username or password not provided or invalid"** 
2. **"password must be at least 8 characters long"** 
3. **"password must be at most 64 characters long"** 
4. **"password must be at most 72 bytes long"** 
5. **"username or password contains invalid characters"**
6. **"unable to make request to pwnedpasswords"** 
7. **"unable to read response from pwnedpasswords"**
8. **"password is leaked"**


Any other error message can be considered a bug.

## Security and Reliability
 - The endpoint is secure and does not pose inherent risks by its design.
 - There is not an admin panel for this solution. 
 - Our deployment scales horizontally in a serverless environment, providing robustness against DDoS attacks targeting the request endpoint.
 - Minor spelling/grammatic errors are not bugs unless they prevent the usage of the platform. 

## Resources
 - Timeout: 8000ms
 - Memory: 256mb

## Prizes
Payed using Bitcoin Lightning Network:
 - Issues: 5.000 SATS (R$ 10) per issue.
 - Bugs: 50.000 SATS (R$ 100) per bug.
 - Critical bugs: 100.000 (R$ 200) per critical bug. 

