package service

import (
	"context"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	db "github.com/AkezhanOb1/customer/repository"
)

func (c *Customer) GetCustomerByEmail(ctx context.Context, request *pb.GetCustomerByEmailRequest) (*pb.GetCustomerByEmailResponse, error) {

	customer, err := db.GetCustomerByEmailRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

