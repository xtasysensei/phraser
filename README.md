
# Phraser

**Phraser** is a CLI tool designed to securely store and manage the seed phrases for your cryptocurrency wallets. 
Whether you're a seasoned crypto enthusiast or just getting started, Phraser simplifies the management of multiple 
wallet seed phrases in a secure and organized way.

## Table of Contents

- [Synopsis](#synopsis)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Options](#options)
- [Examples](#examples)
- [License](#license)

## Synopsis

Phraser allows you to save and manage the seed phrases for your crypto wallets efficiently. By using Phraser, you can 
create a secure store, add wallets with their corresponding seed phrases, and retrieve them when needed.

## Installation
To install Phraser, you can either download the binary or build it from source.

### Download Binary
Head over to the [releases page]() and download the version suitable for your operating system.

### Build from Source
Ensure you have Go installed on your system. Then, clone the repository and build the tool:

```sh
git clone https://github.com/xtasysensei/phraser.git
cd phraser
make build
```
After building, you can move the phraser binary to a directory in your PATH for easy access.

## Usage

Phraser is easy to use with a set of straightforward commands. Here's a basic overview:

```sh
   phraser [command] [flags]
```
## Commands

- **completion**: Generate the autocompletion script for the specified shell.
- **create**: Creates a wallet in your store.
- **gendoc**: Generate Markdown documentation for all commands.
- **get**: Retrieves data stored in a wallet.
- **init**: Initializes a store.

## Options

```sh
  -a, --amount int      amount of phrases to be inputted
  -h, --help            help for phraser
  -s, --store string    name of the store to access
  -t, --toggle          Help message for toggle
  -w, --wallet string   name of the wallet to be created
```

## Examples

1. **Initialize a new store:**
   ```sh
   phraser init --store myStore
   ```

2. **Create a new wallet:**
   ```sh
   phraser create --store myStore --wallet myWallet --amount 12
   ```

3. **Retrieve a wallet's seed phrase:**
   ```sh
   phraser get --store --wallet myWallet
   ```
## More info
For more comprehensive info, refer to the docs directory
## License

This project is licensed under the MIT License.

