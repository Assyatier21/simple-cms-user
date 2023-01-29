# Simple CMS User

Welcome to the Simple CMS User Service. An open-source Content Management System based on the echo framework. As users, we can use the features provided by this service in the form of searching for articles and categories including details of each item. This service has implemented clean architecture principles, a practical software architecture solution from Robert C. Martin (known as Uncle Bob).

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

- Clone the git repository:

```
$ git clone https://github.com/Assyatier21/simple-cms-user.git
$ cd simple-cms-user
```

- Install Dependencies

```
$ go mod tidy
```

- Create `config` folder in root path, then create a file `connection.go` in that folder containing this following code:

```
package config

const (
	User     = "YOUR_USERNAME_HERE"
	Password = "YOUR_PASSWORD_HERE"
	Host     = "localhost"
	Port     = "5432"
	Database = "YOUR_DATABASE_HERE"
	Schema   = "YOUR_SCHEMA_HERE"
	Sslmode  = "disable"
)
```

alternatively, we can just run this following command using makefile:

```
$ make all
```

### Running

```
$ go run cmd/main.go
```

### Features

This service has the following API endpoints:

- `/v1/articles`: get list of articles use limit and offset
- `/v1/article`: get article details by id
- `/v1/categories`: get list of categories
- `/v1/category`: get category details by id

We can test the endpoint using the postman collection in `simple-cms-user/tools`.

### Testing

```
$ go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-cms-user/blob/master/LICENSE) file for details.
