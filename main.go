package main

import (
	"fmt"
	categoryService "my-grpc/m/application/service"
	"my-grpc/m/infra/database"
	"my-grpc/m/infra/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := database.CategoryDB{}

	server := grpc.NewServer()

	reflection.Register(server)

	pb.RegisterCategoryServiceServer(server, categoryService.NewCategoryService(db))

	listner, err := net.Listen("tcp", ":50051")
	fmt.Println("Listing ", listner.Addr())
	if err != nil {
		panic(err)
	}
	if err := server.Serve(listner); err != nil {
		panic(err)
	}
	fmt.Println("Stopping server")
}
