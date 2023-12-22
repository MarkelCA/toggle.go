# toggles
Minimalistic feature flags sofware made in go.

# Install
```bash
cp .env.sample .env # Modify the parameters if you want
docker compose up --build
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
# Database
Test commands, just for the record
## Redis

## Mongo
Test commands
```bash
mongosh --port 27018
show dbs
use toggles
db.flags.insert({"name":"new-login-page","value":true})
db.flags.findOne({"name":"new-login-page"}).value
```
