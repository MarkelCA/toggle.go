# cli
A cli tool to interact with the toggles api.
## Usage
```
$ ./bin/tg -h
his tool offers utilities to interact with the togggles API.

Usage:
  tg [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  db          Database utilities
  flag        Flag utilities
  help        Help about any command
  user        User utilities

Flags:
  -h, --help   help for tg

Use "tg [command] --help" for more information about a command.
```
Get flags (with pretty print):
```
./bin/tg flag get -p
[
  {
    "name": "fizz",
    "value": true
  },
  {
    "name": "foo",
    "value": true
  }
]
```
Add permissions to a user:
```
$ ./bin/tg user permission add test update_flag
Permission added
```
Initialize the database:
```
$ ./bin/tg db init
Database initialized
```
