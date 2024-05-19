# Go TCP Checker

Go TCP Checker is a lightweight, concurrent TCP connection tester written in Go, designed to quickly verify the availability of TCP services.
## Features

- **Concurrent Checking**:  Leverages Go's concurrency model to check multiple services simultaneously, making efficient use of resources and reducing the time required to check multiple endpoints.
- **Configurable via YAML**: Easy to configure through a YAML file where you can specify services, timeout, concurrency level, and output preferences.
- **Context Support**: Utilizes context for managing cancellation and timeouts, enhancing its responsiveness to user interruption and network conditions.
- **Flexible Output**: Supports outputting results in both text and JSON formats, either to the console or a specified file.
- **Interface-driven Design**: Employs Go interfaces to abstract the service checking mechanism, allowing for easy extension to support additional protocols beyond TCP.

## Getting Started

### Prerequisites

- Go 1.17 or later

### Installation

1. Clone the repository to your local machine:

```sh
git clone https://github.com/syncorenize/go-tcp-checker.git
cd go-tcp-checker/cmd/
```

2. Build the project:

```sh
go build -o tcpchecker .
```

### Configuration

**Customize** the behavior of Go TCP Checker by editing the `config.yaml` file. Here's what each option means:

- `services`: A list of service addresses you want to check. Each service should include the domain name or IP followed by the port number.
- `timeout`: The maximum time (in seconds) to wait for a response from each service.
- `concurrencyLevel`: How many services to check in parallel. Increase this to speed up checks for a large number of services.
- `outputFormat`: Choose "text" for human-readable output or "json" for machine-parsable output, useful for integration with other tools.
- `outputPath`: Specify a file path to save the output. Leave empty to print directly to the console.

As an example:

```yaml
services:
  - address: "example.com:80"
  - address: "127.0.0.1:80"
timeout: 5
concurrencyLevel: 2
outputFormat: "text" // "text" for .txt or "json" for .json
outputPath: "" // leave empty to print results to the console
```
### Usage

Run the Go TCP Checker with the following command:

```sh
./tcpchecker -config path/to/your/config.yaml
```
**Replace** `path/to/your/config.yaml` with the actual path to your configuration file.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any bugs or feature requests.

## License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

