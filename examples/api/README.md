# api
A REST API example to interact with the toggles api.

## Usage
Using the `curl` utility. Each one of the examples consist on a command followed by the response given by the REST server. This example uses a self-signed certificate for the TLS encryption, so the `-k` flag is used to skip the certificate validation.

*The output of the following commands could not be the same as the one shown here, it'll depend on the flags stored on your database. The server URL will also depend on four dotenv configurations. Check the [configuration docs](https://github.com/MarkelCA/toggles/blob/main/README.md#configure) for more info*
### Authentication
Most of the endpoints are authenticated, so you'll need to login first. The authentication mechanism is JWT, so you'll need to provide a valid username and password to get a token, as well as a valid api key. Then the token will be used to authenticate the requests.

To check your api keys you can run the following query with the [mongo shell](https://www.mongodb.com/docs/mongodb-shell/):
This example uses the default port (27018) for the mongo shell, but you can change it to the one you're using.
```bash
mongosh --port 27018 --eval 'db.users.find()' toggles
```
Now you can use the api key to logi in and get a token:
```bash
curl -k -X POST -H "Content-Type: application/json" -H "X-Api-key:<api-key>" -d '{"username":"admin", "password":"admin"}' https://localhost:3000/login
```

### Features
Create a flag:
```bash
curl -k https://localhost:3000/flags --data '{"name":"flag2","value":false}' -H "Authorization: Bearer <token id>" -H "X-Api-key:<api-key>"
```
Get all flags:
```bash
curl -k https://localhost:3000/flags -H "Authorization: Bearer <token id>" -H "X-Api-key:<api-key>"
```
Get a flag value:
```bash
curl -k https://localhost:3000/flags/flag2 -H "Authorization: Bearer <token id>" -H "X-Api-key:<api-key>"
```
Update a flag value:
```bash
curl -k https://localhost:3000/flags/flag2 -X PUT --data '{"value":false}' -H "Authorization: Bearer <token id>" -H "X-Api-key:<api-key>"
```
