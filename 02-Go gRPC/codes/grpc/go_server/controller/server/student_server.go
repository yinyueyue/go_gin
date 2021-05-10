package server

import (
	"golang.org/x/net/context"
	"grpc/go_server/proto/student"
)

type StudentServer struct{}


// rpc服务端实现
func (s *StudentServer) GetById(ctx context.Context, request *student.StudentRequest) (*student.Student, error) {
	return &student.Student{Id: request.Id, Name: "grpc", Age: 30}, nil

}

