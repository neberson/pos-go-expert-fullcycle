package main

import (
	"database/sql"
	"net"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/gRPC/internal/database"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/gRPC/internal/pb"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
