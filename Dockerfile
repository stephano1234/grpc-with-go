FROM golang:1.19

RUN apt update
RUN apt upgrade -y
RUN apt install sqlite3 -y

WORKDIR /usr/src/grpc-go

EXPOSE 50051

# docker network create grpc-go-net
# docker build -t stephano1234/grpc-go .
# docker run --rm -v "$(pwd)"/:/usr/src/grpc-go stephano1234/grpc-go go mod init github.com/stephano1234/grpc-go
# docker run --rm -v "$(pwd)"/:/usr/src/grpc-go stephano1234/grpc-go go mod tidy
# docker run --rm --network grpc-go-net --name grpc-go -it -v "$(pwd)"/:/usr/src/grpc-go stephano1234/grpc-go go run cmd/server/server.go
# docker run --rm --network grpc-go-net -it ghcr.io/ktr0731/evans:latest repl -r --host grpc-go --port 50051