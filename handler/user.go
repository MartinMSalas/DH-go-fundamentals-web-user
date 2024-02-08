package handler

import (
	"DH-go-fundamentals-web-user/internal/user"
	"DH-go-fundamentals-web-user/transport"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)


func NewUserHTTPServer(ctx context, router *http.ServeMux, router *http.ServeMux, endpoints user.Endpoints) {
	server.HandleFunc("/users", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tran := transport.New(w, r, ctx)

		switch r.Method {
		case http.MethodGet:
			tran.Server(
				transport.Endpoint(endpoints.GetAll),
				decodeGetAllUser,
				encodeResponse,

			)
		case http.MethodPost:
			endpoints.Create(w, r)
		default:
			invalidMethod(w)
		}
	}
}
func decodeGetAllUser(ctx context.Context, r *http.Request) (request interface{}, err error){
	return nil, nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	status := http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, data)
}
fmt.Println("MakeEndpoints")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MakeEndpoints 2")
		switch r.Method {
		case http.MethodGet:
			GetAllUser(ctx, s, w)
		case http.MethodPost:
			//PostUser(ctx, s, w, r)
			decode := json.NewDecoder(r.Body)
			var req CreateReq
			if err := decode.Decode(&req); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, req)
		default:
			invalidMethod(w)
		}
	}