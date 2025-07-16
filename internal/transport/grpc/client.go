package grpc

import (
	userpb "github.com/Brrocat/project-protos/proto/user"
	"google.golang.org/grpc"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	return userpb.NewUserServiceClient(conn), conn, nil
}
