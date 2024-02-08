package user

import (
	domain "DH-go-fundamentals-web-user/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	
	"net/http"
	"strings"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

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

func MakeEndpoints(ctx context.Context, s Service) Controller {
	fmt.Println("MakeEndpoints")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MakeEndpoints 2")
		switch r.Method {
		case http.MethodGet:
			GetAllUser(ctx, s, w)
		case http.MethodPost:
			//PostUser(ctx, s, w, r)
			decode := json.NewDecoder(r.Body)
			var u domain.User
			if err := decode.Decode(&u); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, u)
		default:
			invalidMethod(w)
		}
	}
}



func invalidMethod(w http.ResponseWriter) {
	status := http.StatusMethodNotAllowed
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "method not allowed"}`, status)
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
func PostUser(ctx context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(CreateReq)

	if strings.TrimSpace(req.FirstName) == "" {
		MsgResponse(w, http.StatusBadRequest, "Invalid first name")
		return
	}
	if strings.TrimSpace(req.LastName) == "" {
		MsgResponse(w, http.StatusBadRequest, "Invalid last name")
		return
	}
	if strings.TrimSpace(req.Email) == "" {
		MsgResponse(w, http.StatusBadRequest, "Invalid email")
		return
	}
	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	DataResponse(w, http.StatusCreated, user)
}
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
