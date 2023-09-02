# PayWalletEngine (WIP)
[![Programming Language](https://img.shields.io/badge/Language-Go-success?style=flat-square)](https://go.dev)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-success.svg?style=flat-square)](https://github.com/Goodnessuc/Go-linker/pulls)

## Introduction

This is a basic MVP version of a bank account system implemented in Go using the hexagonal architecture. It includes features such as JWT authentication, account management, transaction processing, and user management.

## Table of Contents

- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [API Documentation](docs/API.md)
- [Built With](#built-with)
- [Contributing](#contributing)
- [License](#license)


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.15 or later)
- Make

### Installation

1. Clone the repository

```
git clone https://github.com/goodnessuc/paywalletengine.git
```

2. Change directory to the project folder

```
cd paywalletengine
```

3. Install dependencies

```
make install
```

4. Build the application

```
make build
```

5. Run the application

```
./cmd/server/main.go
```

## Usage

To start the server, run the following command:

```
./cmd/server/main.go
```

The server will start and listen on the default port 8080.

## Testing

To run the tests, run the following command:

```
make test
```

## Technologies and Tools

- [Go](https://golang.org/) - The programming language used
- [GORM](https://www.gorm.io/gorm) - Database ORM tool
- [The Gorilla Mux Router](https://www.github.com/gorilla/mux) - Go based HTTP routing tool
- [JWT](https://jwt.io/) - JSON Web Tokens for authentication
- [Make](https://www.gnu.org/software/make/) - Build automation tool


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

