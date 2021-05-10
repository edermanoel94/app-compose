# App Compose

Run many services with one command.

## Examples

```json
# some_config.json
[
  {
    "name": "ping",
    "entrypoint": {
      "exec": "ping",
      "args": ["8.8.8.8"]
    }
  },
  {
    "name": "check files",
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

## Build from source

```sh
$ git clone https://github.com/edermanoel94/app-compose
$ go install
```

# TODO List

- [ ] Working with service daemon to cancel if some error occurs

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

## License

This project is licensed under the terms of the MIT license.
