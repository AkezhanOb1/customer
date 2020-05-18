package client

import (
	"context"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	"github.com/AkezhanOb1/customer/config"
	"google.golang.org/grpc"

)



//CheckOwnerPassword is
func CheckOwnerPassword(ctx context.Context, request *pb.CheckCustomerPasswordRequest) (*pb.CheckCustomerPasswordResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewCustomerServiceClient(cc)

	valid, err := c.CheckCustomerPassword(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return valid, nil
}
