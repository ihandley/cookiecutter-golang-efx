package router

import (
	"log"
	"midigator-portfolios/cookiecutter-golang/api/grpc/todo"
	"midigator-portfolios/cookiecutter-golang/app/initializer"
	"net"

	"google.golang.org/grpc"
)

// Init will initialize grpc router
func Init(dependencies initializer.Services, port string) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	t := todo.Init(dependencies)

	todo.RegisterTodoServiceServer(grpcServer, t)

	return grpcServer.Serve(listen)
}
