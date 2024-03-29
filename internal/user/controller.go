package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		GetAll Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	fmt.Println("MakeEndpoints")
	return Endpoints{
		Create: makeCreateEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
	}

}

func makeGetAllEndpoint(s Service) Controller {
	fmt.Println("maketGetAllEndpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}
func makeCreateEndpoint(s Service) Controller {
	fmt.Println("makeCreateEndpoint")

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("makeCreateEndpoint in return")
		req := request.(CreateReq)

		if strings.TrimSpace(req.FirstName) == "" {
			return nil, errors.New("first name is required")

		}
		if strings.TrimSpace(req.LastName) == "" {

			return nil, errors.New("last name is required")
		}
		if strings.TrimSpace(req.Email) == "" {

			return nil, errors.New("email is required")
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {

			return nil, err
		}

		return user, nil
	}
}


/*
	func PostUser(w http.ResponseWriter, u User) {
		users = append(users, u)
		DataResponse(w, http.StatusCreated, u)
	}
*/
