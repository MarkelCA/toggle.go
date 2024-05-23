# api
A REST API example to interact with the toggles api.

## Usage
Using the `curl` utility. Each one of the examples consist on a command followed by the response given by the REST server.

*The output of the following commands could not be the same as the one shown here, it'll depend on the flags stored on your database. The server URL will also depend on four dotenv configurations. Check the [configuration docs](https://github.com/MarkelCA/toggles/blob/main/README.md#configure) for more info*

Create a flag:
```bash
curl localhost:3000/flags --data '{"name":"flag2","value":false}'
```
Get all flags:
```bash
curl localhost:3000/flags
```
Get a flag value:
```bash
curl localhost:3000/flags/flag2
```
Update a flag value:
```bash
curl localhost:3000/flags/flag2 -X PUT --data '{"value":false}'
```
