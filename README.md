# API Rest con Go (Golang) y PostgreSQL

Development of a simple Rest API with [Go](https://golang.org/) and the [PostgreSQL](https://www.postgresql.org/) database engine.

![PostgreSQL](https://img.shields.io/badge/PostgreSQL-9.6-lightblue.svg?logo=postgresql&longCache=true&style=flat) ![Go](https://img.shields.io/badge/Golang-1.13.4-blue.svg?logo=go&longCache=true&style=flat)

## Getting Started

This project uses the **Go** programming language (Golang) and the **PostgreSQL** database engine.

### Prerequisites

[PostgreSQL](https://www.postgresql.org/) is required in version 9.6 or higher and [Go](https://golang.org/) at least in version 1.12

### Installing

The following dependencies are required:

* github.com/go-chi/chi

* github.com/joho/godotenv

* github.com/lib/pq

* golang.org/x/crypto

* github.com/dgrijalva/jwt-go

#### Using GOPATH

```bash
go get github.com/go-chi/chi

go get github.com/joho/godotenv

go get github.com/lib/pq

go get golang.org/x/crypto

go get github.com/dgrijalva/jwt-go
```

#### Using GOMODULE

```bash
go build ./cmd/microblog
```

## Deployment

Clone the repository

```bash
git clone git@github.com:orlmonteverde/go-postgres-microblog.git
```

Enter the repository folder

```bash
cd go-postgres-microblog
```

Build the binary

```bash
go build ./cmd/microblog/
```

Run the program

```bash
# In Unix-like OS
./microblog

# In Windows
microblog.exe
```

### API Documentation

[Swagger](https://app.swaggerhub.com/apis/orlmonteverde/go-postgres-microblog/1.0.0)

## Built With

* [chi](https://github.com/go-chi/chi) - High performance, extensible, minimalist Go web framework

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/orlmonteverde/go-postgres-microblog/tags).

## Authors

* **Orlando Monteverde** - *Initial work* - [orlmonteverde](https://github.com/orlmonteverde)
