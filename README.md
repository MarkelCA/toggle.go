# toggles
Toggles it's a feature flags application made in Go. It uses mongodb as disk storage and redis for flag value's cache management. Counts with implementations as a REST API and as a Websocket server. There's also plans to build SDKs soon.

# Configure
In this step you'll be able to configure the port where the service will run (`APP_PORT`) and the desired implementation (`APP`). The `APP` value should be a folder from `cmd`. Currently there's two options: `api` (rest api) and `ws` (websocket). 
```bash
cp .env.sample .env
```
_Example env file:_
```env
APP=ws
APP_PORT=3000
```

# Install
Using [docker](https://docs.docker.com/desktop/):
```bash
git clone https://github.com/MarkelCA/toggles
cd toggles
docker compose up --build
```
