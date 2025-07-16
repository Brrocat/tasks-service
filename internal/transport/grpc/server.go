package grpc

import (
	taskpb "github.com/Brrocat/project-protos/proto/task"
	"log"
	"net"

	userpb "github.com/Brrocat/project-protos/proto/user"
	"github.com/Brrocat/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(srv, handler)

	log.Println("Starting Tasks gRPC server on :50052")
	return srv.Serve(lis)
}
