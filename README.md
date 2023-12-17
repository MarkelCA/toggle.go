# toggles
Minimalistic feature flags sofware made in go.

# Install
```bash
docker compose up
```

# Examples
```bash
# List all flags
curl localhost:8080/flags -v

# Create flag
curl localhost:8080/flags -d '{"name":"new-login-page","value":true}' -v

# Get single flag
curl localhost:8080/flags/new-login-page -v

# Update flag
curl localhost:8080/flags/new-login-page -X PUT --data '{"value":false}'
```
