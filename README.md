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

1. **"username or password not provided or invalid"** 
2. **"password must be at least 8 characters long"** 
3. **"password must be at most 64 characters long"** 
4. **"password must be at most 72 bytes long"** 
5. **"username or password contains invalid characters"**
6. **"unable to make request to pwnedpasswords"** 
7. **"unable to read response from pwnedpasswords"**
8. **"password is leaked"**


Any other error message can be considered a bug.
