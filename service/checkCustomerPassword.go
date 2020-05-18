package service

import (
	"context"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	db "github.com/AkezhanOb1/customer/repository"
)

func (c *Customer) CheckCustomerPassword(ctx context.Context, request *pb.CheckCustomerPasswordRequest) (*pb.CheckCustomerPasswordResponse, error) {
	password, err := db.GetCustomerPasswordRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	err = comparePassword(*password, request.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.CheckCustomerPasswordResponse{
		Valid:true,
	}, nil
}

