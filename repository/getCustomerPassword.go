package repository

import (
	"context"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	"github.com/AkezhanOb1/customer/config"
	"github.com/jackc/pgx/v4"

)

//GetCustomerPasswordRepository is a
func GetCustomerPasswordRepository (ctx context.Context, request *pb.CheckCustomerPasswordRequest) (*string, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT password FROM customer WHERE email=$1;`

	var password string
	err = conn.QueryRow(ctx, sqlQuery, request.GetEmail()).Scan(&password)
	if err != nil {
		return nil, err
	}

	return &password, nil

}
