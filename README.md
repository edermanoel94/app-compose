# App Compose

Run many services with one command.

## Examples

```json
[
  {
    "name": "ping",
    "path": "",
    "entrypoint": {
      "exec": "ping",
      "args": ["8.8.8.8"]
    }
  },
  {
    "name": "check files",
    "path": "",
    "entrypoint": {
      "exec": "ls",
      "args": ["-la"]
    }
  }
]
```

```sh
$ app-compose run some_config.json
```

## Prerequisites

- Go 1.16

## Installation

We using go modules and get always the last version

```
$ go get github.com/edermanoel94/app-compose@latest
```

# TODO List

- [ ] Working with service daemon to cancel if some error occurs

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

## License

This project is licensed under the terms of the MIT license.
