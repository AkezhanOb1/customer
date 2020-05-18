package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AkezhanOb1/customer/api/graphql/graph/generated"
	"github.com/AkezhanOb1/customer/api/graphql/graph/model"
	pba "github.com/AkezhanOb1/customer/api/proto/auth"
	pb "github.com/AkezhanOb1/customer/api/proto/customer"
	"github.com/AkezhanOb1/customer/client"
	"github.com/AkezhanOb1/customer/pkg"
)

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	var req = &pb.CreateCustomerRequest{
		CustomerFirstName:         input.CustomerFirstName,
		CustomerSecondName:        input.CustomerSecondName,
		CustomerEmail:             input.CustomerEmail,
		CustomerPhoneNumberPrefix: input.CustomerPhoneNumberPrefix,
		CustomerPhoneNumber:       input.CustomerPhoneNumber,
		CustomerPassword:          input.CustomerPassword,
	}

	customer, err := client.CreateCustomer(ctx, req)

	if err != nil {
		return nil, nil
	}

	b, err := pkg.Serializer(customer)
	if err != nil {
		return nil, err
	}

	var response *model.CreateCustomerResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	var tokenRequest = pba.GenerateCustomerTokenRequest{
		Credentials: &pba.ClientCredentials{
			Email:    customer.Customer.GetCustomerEmail(),
			Password: "",
		},
	}
	token, err := client.GenerateToken(context.Background(), &tokenRequest)
	if err != nil {
		return nil, err
	}

	response.Token = &model.Token{
		AccessToken:  token.Token.GetAccessToken(),
		RefreshToken: token.Token.GetRefreshToken(),
		ExpiresIn:    token.Token.GetExpiresIn(),
		TokenType:    token.Token.GetTokenType(),
	}

	return response, nil
}

func (r *mutationResolver) CreateCustomerToken(ctx context.Context, input model.CreateCustomerTokenRequest) (*model.CreateCustomerTokenResponse, error) {
	var checkPasswordRequest = pb.CheckCustomerPasswordRequest{
		Email:    input.Email,
		Password: input.Password,
	}
	checkResult, err := client.CheckOwnerPassword(context.Background(), &checkPasswordRequest)
	if err != nil {
		return nil, err
	}

	var tokenRequest = pba.GenerateCustomerTokenRequest{
		Credentials: &pba.ClientCredentials{
			Email:    input.Email,
			Password: "",
		},
	}

	token, err := client.GenerateToken(context.Background(), &tokenRequest)
	if err != nil {
		return nil, err
	}

	if checkResult.GetValid() != true {
		return nil, fmt.Errorf("email or password is not correct")
	}

	return &model.CreateCustomerTokenResponse{
		Token: &model.Token{
			AccessToken:  token.Token.GetAccessToken(),
			RefreshToken: token.Token.GetRefreshToken(),
			ExpiresIn:    token.Token.GetExpiresIn(),
			TokenType:    token.Token.GetTokenType(),
		},
	}, nil
}

func (r *queryResolver) GetCustomerByEmail(ctx context.Context, input model.GetCustomerByEmailRequest) (*model.GetCustomerByEmailResponse, error) {
	var req = &pb.GetCustomerByEmailRequest{
		Email: input.Email,
	}

	customer, err := client.GetCustomerByEmail(ctx, req)

	if err != nil {
		return nil, nil
	}

	b, err := pkg.Serializer(customer)
	if err != nil {
		return nil, err
	}

	var response *model.GetCustomerByEmailResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *queryResolver) GetCustomerTokenInfo(ctx context.Context, input model.GetCustomerTokenInfoRequest) (*model.GetCustomerTokenInfoResponse, error) {
	var tokenInfoRequest = pba.RetrieveCustomerTokenInformationRequest{
		AccessToken: input.AccessToken,
	}
	info, err := client.RetrieveTokenInformation(context.Background(), &tokenInfoRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetCustomerTokenInfoResponse{
		Email:     info.GetEmail(),
		IssuedAt:  info.GetIssuedAt(),
		ExpiresAt: info.GetExpiresAt(),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
