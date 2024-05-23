# websocket
A websocket example connection to the toggles api.

## Usage
Using the [`websocat`](https://github.com/vi/websocat) utility. Each one of the examples consist on a command followed by the response given by the websocket server.

Connect to websocket + get a flag value:
```
websocat ws://127.0.0.1:3000/ws
```
Get a flag value (non existent):
```
{"command":"get","data":"myflag"}
{"status":404,"value":"toggles: Flag not found"}

```
Create a flag:
```
{"command":"create","data":{"name":"myflag","value":false}}
{"status":201,"value":{"name":"myflag","value":false}}

```
Get the flag value (after creation):
```
{"command":"get","data":"myflag"}
{"status":200,"value":false}
```
Update the flag value:
```
{"command":"update","data":{"name":"myflag","value":true}}
{"status":201,"value":null}
```
