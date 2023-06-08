# Gandi Alias CLI

A command line tool written in Go to manage Gandi Aliases.

## Installation

To install the Gandi Alias CLI, run the following command:

    go get github.com/aegypius/gandi-alias

## Usage

To use the Gandi Alias CLI, run the following command:

    gandi-alias [command] [options]

### Commands

- `add [alias] [destination]`: Add a new alias to your Gandi account.
- `list`: List all aliases in your Gandi account.
- `remove [alias]`: Remove an alias from your Gandi account.
- `help`: Show help information for the Gandi Alias CLI.

### Options

- `-k, --api-key [key]`: Set the API key to use for authentication. If not specified, the API key will be read from the `GANDI_API_KEY` environment variable.
- `-v, --verbose`: Enable verbose output.

## Contributing

Contributions are welcome! If you have any bug reports, feature requests, or patches, please submit them via pull request or issue on GitHub.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
