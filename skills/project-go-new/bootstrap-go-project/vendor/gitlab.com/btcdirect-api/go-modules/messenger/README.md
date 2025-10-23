# Access private package

If you did this before, skip steps 1-4.

1. Generate a Gitlab access token with the `api` scope enabled
2. Create a ~/.netrc file:
```bash
machine gitlab.com
    login <your gitlab username>
    password <the token created in step 1>
```
3. `chmod 600 ~/.netrc`
4. `export GOPRIVATE=gitlab.com/btcdirect-api/*`
5. `go get gitlab.com/btcdirect-api/go-modules/messenger`
