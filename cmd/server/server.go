package main

import (
	"net"
	"database/sql"

	"github.com/stephano1234/grpc-go/internal/database"
	"github.com/stephano1234/grpc-go/internal/service"
	"github.com/stephano1234/grpc-go/internal/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3" 
)

func main() {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	categoryService := service.NewCategoryService(database.NewCategory(db))
	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryService)
	reflection.Register(server)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}