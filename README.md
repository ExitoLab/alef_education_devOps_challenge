### Alef Education DevOps Challenge

The Application for this challenge is written in Go Programming language using Gin framework. The workflow for running this challenge is written using a Makefile

The application is created using: 

1. Go - Gin Framework 
2. Mongodb - database

## How to run the app. 
1. Using docker-compose 
2. Using helm charts on Kubernetes 

## Running the application using a makefile with docker-compose

1. Command to run / start the application. Run `make compose-up`
2. To stop the app using docker-compose. Run `make compose-down`

## Running the application using a makefile on kubernetes / helm

1. Command to start the application using helm. Run `make install-api`



You can run the app using the old fashion way.

## To deploy 

1. Create mongodb 
2. Install golang 
3. Run command `go build -o new -v`
4. Run command `./new` 
5. To run in background use this command `nohup ./new &`


## Kill running app on port 8000

To kill the app while debugging run this command `kill -9 $(lsof -t -i:8000)`

## More Debugging 

To debug use `netstat -plten`You can run the app using docker-compose or the old fashion way.

## To deploy 

1. Create mongodb 
2. Install golang 
3. Run command `go build -o new -v`
4. Run command `./new` 
5. To run in background use this command `nohup ./new &`


## Kill running app on port 8000

To kill the app while debugging run this command `kill -9 $(lsof -t -i:5000)`

## More Debugging 

To debug use `netstat -plten`