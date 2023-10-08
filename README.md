# PayWalletEngine

<p>
   <a href="http://makeapullrequest.com"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat" alt=""></a>
   <a href="https://golang.org"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="made-with-Go"></a>
   <a href="https://goreportcard.com/report/github.com/goodnessuc/paywalletengine"><img src="https://goreportcard.com/badge/github.com/goodnessuc/paywalletengine" alt="GoReportCard"></a>
   <a href="https://github.com/goodnessuc/paywalletengine"><img src="https://img.shields.io/github/go-mod/go-version/goodnessuc/paywalletengine.svg" alt="Go.mod version"></a>
   <a href="https://github.com/goodnessuc/paywalletengine/blob/master/LICENSE"><img src="https://img.shields.io/github/license/goodnessuc/paywalletengine.svg" alt="LICENSE"></a>
</p>

## Introduction

This is a basic MVP version of a bank account system implemented in Go using the hexagonal architecture. It includes
features such as JWT authentication, account management, transaction processing, and user management.

## Table of Contents

- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#Installation-and-Setup)
- [Usage](#usage)
- [Docker Deployment](#docker-deployment)
    - [Prerequisites](#prerequisites-1)
    - [Running with Docker Compose](#running-with-docker-compose)
- [Testing](#testing)
- [API Documentation](docs/readme.md)
- [Technologies and Tools](#technologies-and-tools)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

- Go (version 1.15 or later)
- Make

### Installation and Setup

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

## Docker Deployment

This application can be run using Docker Compose, which sets up the required services including a database and the
application itself.

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Running with Docker Compose

1. Ensure that Docker and Docker Compose are installed on your machine.

2. Navigate to the project directory:

```
cd paywalletengine
```

3. Build and run the application and its services using Docker Compose:

```
docker-compose up --build
```

This will start up the `db` service using the PostgreSQL image and the `api` service which is your application. The
application will connect to the database, and both will be accessible from your machine.

To stop the services, simply run:

```
docker-compose down
```

This will gracefully shut down the containers and remove them.

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

Contributions are welcome! Feel free to open a pull request right away.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
