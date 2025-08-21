# WebUpDown

WebUpDown is a simple command-line tool written in Go that checks whether a website is up or down by attempting a TCP connection to the specified domain and port.

## Features

- **Check website availability** via TCP.
- **Customizable port** (default is 80).
- **Clear status output** indicating if the site is reachable.

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/RajaPremSai/healthcheckweb.git
   cd healthcheckweb
   ```

2. **Build the project:**
   ```sh
   go build -o webupdown
   ```

## Usage

```sh
./webupdown --domain example.com [--port 80]
```

Or with short flags:

```sh
./webupdown -d example.com -p 443
```

- `--domain`, `-d`: Domain name to check (**required**)
- `--port`, `-p`: Port number to check (default: 80)

## Example

```sh
./webupdown --domain google.com
```

**Sample Output:**

```
[UP] google.com is reachable, From : 192.168.1.10:54321 to 142.250.72.14:80
```
Or

```
[DOWN] amazo website is not reachable, Error : dial tcp: lookup amazo: no such host
```

## Dependencies

- [urfave/cli/v2](https://github.com/urfave/cli) - To build command-line interface (CLI) applications quickly and cleanly. It provides a structured way to define commands, flags, arguments, and subcommands for CLI tools.

## License

MIT License
