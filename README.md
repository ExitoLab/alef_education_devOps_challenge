### Alef Education DevOps Challenge


## Dockerize 

Docker and docker-compose is in progress. Currently deploying the code fashion way. 

## To deploy 

1. Create mongodb 
2. Install golang 
3. Run command `go build -o new -v`
4. Run command `./new` 
5. To run in background use this command `nohup ./new &`


## Kill running app on port 5000

To kill the app while debugging run this command `kill -9 $(lsof -t -i:5000)`
