# Boolean As a Service 
- Ability to create, get, update and delete a boolean value <br />
- Storing the data in mysql database <br />
- Exposing  RESTful endpoints for each operation <br />
- Handles POST, GET, PATCH and DELETE requests <br />

## Requirements
- mysql 8.0.21 or higher
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
 $ go mod download
```
Run below command for setting ENVIRONMENT variables
```sh
 $ source .env
```
Start the service
```sh
 $ go run main.go
```
## API
```
 base url: http://localhost/ 
```
use __POST /__ method to create a new boolean
__request:__
```
 {
   "value":true,
    "key": "name" // this is optional
 }
```
- value should be either true or false(boolean, not string) <br />
- key should be string <br />

__response:__
```
 {
   "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
   "value": true,
   "key": "name"
 }
```


use __GET /:id__ method to get an existing boolean
__response:__
```
 {
   "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
   "value": true,
   "key": "name"
 }
```
use __PATCH /:id__ method to update an existing boolean
__request:__
```
 {
   "value":false,
   "key": "new name" // this is optional
 }
```
__response:__
```
 {
   "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
   "value": false,
   "key": "new name"
 }
```
use __DELETE /:id__ method to delete an existing boolean
__response:__
```
 HTTP 204 No Content
```
## Container

Go to [boolean-as-a-service](https://hub.docker.com/r/dilipchauhan1998/boolean-as-a-service) for the docker image of the boolean service

