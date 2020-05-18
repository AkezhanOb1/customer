package repository

import (
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	"github.com/AkezhanOb1/customer/config"
	"context"
	"github.com/jackc/pgx/v4"

)





//GetCustomerByEmailRepository is a
func GetCustomerByEmailRepository (ctx context.Context, request *pb.GetCustomerByEmailRequest) (*pb.GetCustomerByEmailResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT id, first_name, second_name, email, phone_number_prefix, phone_number, created_at
	FROM customer WHERE email=$1;`


	var customer pb.Customer
	err = conn.QueryRow(ctx, sqlQuery, request.GetEmail()).Scan(
		&customer.CustomerID,
		&customer.CustomerFirstName,
		&customer.CustomerSecondName,
		&customer.CustomerEmail,
		&customer.CustomerPhoneNumberPrefix,
		&customer.CustomerPhoneNumber,
		&customer.CreatedAt,
	)

	if err != nil {
		return nil, err
	}


	return &pb.GetCustomerByEmailResponse{
		Customer: &customer,
	}, nil
}
