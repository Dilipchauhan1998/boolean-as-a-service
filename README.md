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

## APIs
### POST /
#### request:
```
{
  "value":true,
   "key": "name" // this is optional
}
```
#### response:
```
{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```
value should be either true or false(boolean, not string) <br />
key should be string <br />

### GET / :id
#### response:
```
{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```
### PATCH/:id
#### request:
```
{
  "value":false,
  "key": "new name" // this is optional
}
```
#### response:
```
{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": false,
  "key": "new name"
}
```
### DELETE /:id
#### response:
```
HTTP 204 No Content
```
