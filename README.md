# toggles
Minimalistic feature flags sofware made in go.

# Install
```bash
docker compose up
```

# Examples
```bash
curl localhost:8080/flags -v
curl localhost:8080/flags -d '{"name":"new-login-page","value":true}' -v
curl localhost:8080/flags/new-login-page -v
curl localhost:8080/flags/new-login-page -X PUT --data '{"value":false}'
```
