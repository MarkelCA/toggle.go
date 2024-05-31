# cli
A cli tool to interact with the toggles api.
## Usage
```bash
$ ./bin/tg -h
This tool offers utilities to interact with the togggles API.

Usage:
  tg [flags]
  tg [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  db
  help        Help about any command

Flags:
  -h, --help   help for tg

Use "tg [command] --help" for more information about a command.
```

```bash
$ ./bin/tg db get users
{
	"UserName": "admin",
	"FirstName": "",
	"LastName": "",
	"Role": "admin",
	"Password": "$2a$14$CrC0h1NrVe6k6keeTdDN6OmFwMJZdkovlchdvBMECQYLBKFZ3aa0.",
	"ApiKey": "815226dbe48b90b98a4cd80521d2dabde35490175bd59af08cdc09676d8f01d6"
}
{
	"UserName": "test",
	"FirstName": "",
	"LastName": "",
	"Role": "user",
	"Password": "$2a$14$XveBl98ZncVp39gtFIWjOeqs/oEEjq8106N7kPKsic6NtpfITDt32",
	"ApiKey": "eb29f5e59cbdc055f0ebbe90693a7802618aa6bcdc110ecf2ba22bcb07dfac6c"
}
```
