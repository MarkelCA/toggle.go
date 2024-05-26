# api
A REST API example to interact with the toggles api.

## Usage
Using the `curl` utility. Each one of the examples consist on a command followed by the response given by the REST server.

*The output of the following commands could not be the same as the one shown here, it'll depend on the flags stored on your database. The server URL will also depend on four dotenv configurations. Check the [configuration docs](https://github.com/MarkelCA/toggles/blob/main/README.md#configure) for more info*
### Authentication
Most of the endpoints are authenticated, so you'll need to login first. The authentication mechanism is JWT, so you'll need to provide a valid username and password to get a token. Then the token will be used to authenticate the requests.
```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"admin", "password":"admin"}' http://localhost:3000/login
```

### Features
Create a flag:
```bash
curl localhost:3000/flags --data '{"name":"flag2","value":false}' -H "Authorization: Bearer <token id>"
```
Get all flags:
```bash
curl localhost:3000/flags -H "Authorization: Bearer <token id>"
```
Get a flag value:
```bash
curl localhost:3000/flags/flag2 -H "Authorization: Bearer <token id>"
```
Update a flag value:
```bash
curl localhost:3000/flags/flag2 -X PUT --data '{"value":false}' -H "Authorization: Bearer <token id>"
```
