package service

import (
	"context"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	db "github.com/AkezhanOb1/customer/repository"
)

func (c *Customer) CreateCustomer(ctx context.Context, request *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	hashedPassword, err := hashPassword(request.GetCustomerPassword())
	if err != nil {
		return nil, err
	}
	request.CustomerPassword = hashedPassword

	customer, err := db.CreateCustomerRepository(ctx, request)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
