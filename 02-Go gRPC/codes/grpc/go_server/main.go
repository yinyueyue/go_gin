package main

import (
	"google.golang.org/grpc"
	"grpc/go_server/controller/server"
	"grpc/go_server/proto/hello"
	"grpc/go_server/proto/student"
	"log"
	"net"
)

const (
	Address = "0.0.0.0:9090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// 服务注册
	hello.RegisterHelloServer(s, &server.HelloController{})
	student.RegisterStudentServiceServer(s,&server.StudentServer{})

	log.Println("Listen on " + Address)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
