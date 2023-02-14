## Get running
Install golang (version 1.19) and sqlite3 and run the command:
```
go run cmd/server/server.go
```
Or use the docker to build an image from Dockerfile:
```
docker build -t stephano1234/grpc-go .
```
Create a network for the container (useful for further testing):
```
docker network create grpc-go-net
```
Run the server container:
```
docker run --rm --network grpc-go-net --name grpc-go -it -v "$(pwd)"/:/usr/src/grpc-go stephano1234/grpc-go go run cmd/server/server.go
```
## Test it
Install evans ([see the documentation](https://github.com/ktr0731/evans)) and run:
```
evans repl -r
```
Or use their docker image:
```
docker run --rm --network grpc-go-net -it ghcr.io/ktr0731/evans:latest repl -r --host grpc-go --port 50051
```