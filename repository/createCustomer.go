package repository

import (
	"context"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	"github.com/AkezhanOb1/customer/config"
	"github.com/jackc/pgx/v4"
	"log"

	"strings"
	"time"
)

//CreateCustomerRepository is
func CreateCustomerRepository(ctx context.Context, request *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())


	sqlQuery := `INSERT INTO customer (first_name, second_name, email, password, 
								phone_number_prefix, phone_number, created_at)
				 VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	var customerID int64
	firstName := request.GetCustomerFirstName()
	secondName := request.GetCustomerSecondName()
	email := strings.ToLower(request.GetCustomerEmail())
	password := request.GetCustomerPassword()
	phoneNumberPrefix := request.GetCustomerPhoneNumberPrefix()
	phoneNumber := request.GetCustomerPhoneNumber()
	createdAt := time.Now().Format(time.RFC3339)


	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		firstName,
		secondName,
		email,
		password,
		phoneNumberPrefix,
		phoneNumber,
		createdAt,
	)

	err = row.Scan(&customerID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateCustomerResponse{
		Customer: &pb.Customer{
			CustomerID:                customerID,
			CustomerFirstName:         firstName,
			CustomerSecondName:        secondName,
			CustomerEmail:             email,
			CustomerPhoneNumberPrefix: phoneNumberPrefix,
			CustomerPhoneNumber:       phoneNumber,
			CreatedAt: 				   createdAt,
		},
	}, nil

}
