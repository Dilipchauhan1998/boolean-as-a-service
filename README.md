# Boolean As a Service 
- Ability to create, get, update and delete a boolean value <br />
- Storing the data in mysql database <br />
- Exposing  RESTful endpoints for each operation <br />
- Handles POST, GET, PATCH and DELETE requests <br />

## Requirements
- __Mysql 8.0.21 or higher__ <br />
To install mysql on  macOS follow [install mysql](https://flaviocopes.com/mysql-how-to-install/) <br />

- __go 1.15 or higher__ <br />
To install go on macOS follow [install go](https://www.geeksforgeeks.org/how-to-install-golang-on-macos/) <br />

## Configuration

- Create a new user
```
mysql> CREATE USER 'usernamer'@'localhost' IDENTIFIED BY 'password';
```
- Create a database
```
mysql> CREATE DATABASE database_name
```
- Give the created user all privileges on the created database
```
mysql> GRANT ALL PRIVILEGES ON database_name.* TO 'username'@'localhost' identified by 'password';
mysql> FLUSH PRIVILEGES;
```

- Edit the following lines in the files ".env":

```
export MYSQL_DB_USER="username"     
export MYSQL_DB_PASS="password"
export MYSQL_DB_NAME="database_name"
```
## Installation
### Without Docker
- Clone the repository and keep it in the $GOPATH
- Open the terminal and ``` cd ``` to the cloned repository directory
    ```sh
     $ cd paths/boolean-as-a-service
    ```
- Install all required packages
    ```sh
     $ go mode download
    ```
- Run below command for setting database related ENVIRONMENT variables
    ```sh
     $ source .env
    ```
- Build the service 
    ```sh
     $ go build main.go
    ```
- Run the service 
   ```sh
    $ ./main
   ```
   
### With Docker
Go to [boolean](https://hub.docker.com/r/dilipchauhan1998/boolean) for the docker image of the boolean service

__Create docker image__
- Clone the repository and ```cd``` to the cloned repository directory 
    ```sh
    $ cd paths/boolean-as-a-service
    ```
- Run below command to create image
     ```sh
    $ docker build . -t dilipchauhan1998/boolean
    ```
__OR pull image from docker hub__

- Pull the image to local machine using command
    ```sh
    $ docker pull dilipchauhan1998/boolean
    ```
__Start boolean service__     
- Start default instance
    ```sh
    $ docker run --name boolean -p 80:80 dilipchauhan1998/boolean
    ```
- Start boolean service with your own database configuration  
     ```sh
     $ docker run --name boolean -p 80:80 -e MYSQL_USER=dev -e MYSQL_USER_PWD=dev -e MYSQL_USER_DB=userdb dilipchauhan1998/boolean

     ```

## API
__base url__
```
  http://localhost:80
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

