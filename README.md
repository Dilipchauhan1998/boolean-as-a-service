# Boolean As a Service 
- A service which can be used to create, delete and update boolean values <br />
- Provides REST endpoints for each operation <br />
- Handles POST, GET, PATCH and DELETE requests <br />

## Requirements
- Mysql 8.0.21 or higher
- go 1.15 or higher
## Configuration

Edit the following lines in the files ".env": <br />

```
export MYSQL_DB_USER="xxxxxx"     
export MYSQL_DB_PASS="xxxxxx"
export MYSQL_DB_NAME="xxxxxx"
```
## Installation
Clone the repository and keep it in the $GOPATH <br />
Open the terminal and run <br />
```sh
 $ cd path-to-boolean-as-a-service/boolean-as-a-service
```
Install all required packages <br />
```sh
 $ go get
```
Run below command for setting ENVIRONMENT variables
```sh
 $ source .env
```
Start the service
```sh
 $ go run main.go
```
