# Go rest api

A simple REST API server for a Todo app built with Gin framework in Go.

## Set Up the Repository

```bash
$ git clone https://github.com/ingerstep/go-restapi
$ cd go-restapi
```

# Running with Docker

## Description

To simplify the deployment of the application, you can use Docker.

## Installation

Make sure you have Docker and Docker Compose installed.

```bash
$ docker --version
$ docker compose version | docker-compose --version
```
## Environment

```bash
#create .env in root of the project
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
POSTGRES_LOCAL_PORT=
POSTGRES_PORT=
POSTGRES_HOST=
SSL_MODE=
SERVER_PORT=
SERVER_LOCAL_PORT=
```
## Running the Application

### Navigate to the root directory of the project.

Start Docker Compose:

```bash
$ docker-compose up --build
```

### Stopping the Application

To stop the application, use:

```bash
$ docker-compose down
```

- **Swagger will be accessible at http://localhost:${SERVER_PORT}/swagger/index.html.**
