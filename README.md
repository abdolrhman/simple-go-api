# simple-go-api-wajve


![Build](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/workflows/Go/badge.svg)


This is a simple task, to integrate some json data with mongodb and create an API end point using go to get this data
A few things to note in the project:
* **[Github Actions Workflows](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/tree/main/.github/workflows)** - Pre-configured Github Actions to run automated builds and publish image to Github Packages
* **[Dockerfile](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/blob/main/Dockerfile)** - Dockerfile to generate docker builds.
* **[docker-compose](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/blob/main/docker-compose.yml)** - Docker compose script to start service in production mode.
* **[Containerized Mongo for development](#development)** - Starts a local mongo container with data persistence across runs.
* **[Mongo Driver](https://go.mongodb.org/mongo-driver)** - MongoDB supported driver for Go.
* **[Gorilla Mux](https://go.mongodb.org/mongo-driver)** - HTTP request multiplexer.
* **[.env file for configuration](#environment)** - Change server config like app port, mongo url etc
* **[httptest](#testing)** - Utilities for HTTP testing.

## Installation

#### 1. Clone this repo

```
$ git clone git@github.com:abdolrhman/simple-go-api.git
$ cd simple-go-api.git
```

#### 2. Install dependencies

```
$ go mod vendor
```

## Architecture 
- a function run every time the application starts
- fill the database with the records on the josn
  - if no any
- API that grapes the data from mongodb
- The API accepts parameters if given, that creat that construct the query options
- nginx and workflow supported
- integration test supported

## Development

### Start dev server
Starting the dev server also starts MongoDB as a service in a docker container using the compose script at `docker-compose.yml`.

```
$ go run main.go
```
Running the above commands results in 
* 🌏 **API Server** running at `http://localhost:8080`
<!-- * ⚙️**Swagger UI** at `http://localhost:3000/dev/api-docs` -->
* ⛁ **MongoDB** running at `mongodb://localhost:27017/wajve`

## Packaging and Deployment
#### 1. Build and run without Docker

```
$ go build 
```
#### 2. Run Tests

```
$ cd tests
$ go test
```
#### 3. Run with docker

```
$ docker build -t api-server .
$ docker run -t -i -p 8080:8080 api-server
```

#### 4. Run with docker-compose

```
$ docker-compose up
```


---

## Environment
To edit environment variables, create a file with name `.env` and copy the contents from `.env.default` to start with.

| Var Name  | Type  | Default                           | Description  |
|---|---|-----------------------------------|---|
|  PORT | number  | `8080`                            | Port to run the API server on |
|  MONGO_URL | string  | `mongodb://localhost:27017/wajve` | URL for MongoDB |

<!-- ## Logging
The application uses [winston](https://github.com/winstonjs/winston) as the default logger. The configuration file is at `src/logger.ts`.
* All logs are saved in `./logs` directory and at `/logs` in the docker container.
* The `docker-compose` file has a volume attached to container to expose host directory to the container for writing logs.
* Console messages are prettified
* Each line in error log file is a stringified JSON. -->


### Directory Structure

```
+-- controllers
|   +-- triviaController.go
+-- db
|   +-- db.go
|   +-- db.json
+-- handlers
|   +-- config.go
|   +-- logs.go
|   +-- response.go
+-- models
|   +-- models.go
+-- tests
|   +-- api_test.go
+-- routes
|   +-- routes.go
+-- vendor
+-- nginx
|   +-- dev.conf.d
|   |   +-- nginx.conf
+-- .env
+-- .env.default
+-- .gitignore
+-- docker-compose.yml
+-- Dockerfile
+-- go.mod
+-- go.sum
+-- main.go
+-- README.md
```
