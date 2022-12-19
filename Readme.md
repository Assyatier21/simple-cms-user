# Simple CMS API Documentation

Welcome to the API documentation for the simple CMS Service. This API allows you to get an article and category as user. This service using echo framework as well.


## Endpoints

The API has the following endpoints:
- `/v1/articles`: get list of articles
- `/v1/article`: get details of article by id
- `/v1/categories`: get list of categories
- `/v1/category`: get details of category by id

## Installation and How to run it locally
1. Clone the repository
2. Import the database (PostgreSQL) into your local. SQL file location: `simple-cms-user/tools` 
3.  Then run this command below:
```bash
$ go mod tidy
$ go run cmd/main.go
```
4. We can test the endpoint using the collection located in : `simple-cms-user/tools` 
 
