# Go Api Template

## How to run

### Docker
Download this repository and run
````shell
docker-compose -f docker\docker-compose.yaml up --build
````

### With local go intallation
````shell
go mod vendor
go mod download
go run entrypoints/api/main.go
````
