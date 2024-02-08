package user

import (
	//domain "DH-go-fundamentals-web-user/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
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
		GetAll: maketGetAllEndpoint(s),
	}	
	
}




func makeCreateEndpoint(s 	Service) Controller{
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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

func maketGetAllEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}
func GetAllUser(ctx context.Context, s Service, w http.ResponseWriter) {
	fmt.Println("GetAllUser in Controller")
	users, err := s.GetAll(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	DataResponse(w, http.StatusOK, users)

}

/*
	func PostUser(w http.ResponseWriter, u User) {
		users = append(users, u)
		DataResponse(w, http.StatusCreated, u)
	}
*/

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)
}
func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())

		return
	}
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}
