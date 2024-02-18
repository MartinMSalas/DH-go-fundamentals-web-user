package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	Server(
		endpoint Endpoint,
		decode func(ctx context.Context, r *http.Request) (request interface{}, err error),
		encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
		encodeError func(ctx context.Context, err error, w http.ResponseWriter),
		
	)

}

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type transport struct {
	w http.ResponseWriter
	r *http.Request
	ctx context.Context
}

func New(w http.ResponseWriter, r *http.Request,ctx context.Context) Transport {
	return &transport{w: w, r: r, ctx: ctx}


}

func (t *transport) Server(
	endpoint Endpoint,
	decode func(ctx context.Context, r *http.Request) (request interface{}, err error),
	encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	encodeError func(ctx context.Context, err error, w http.ResponseWriter),
) {
	
	data, err := decode(t.ctx, t.r)
	if err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}

	response, err := endpoint(t.ctx, data)
	if err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}

	if err := encode(t.ctx, t.w, response); err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}
}
	
func Clean(url string) ([]string, int){
	if url[0] != '/' {
		url = "/" + url
	}

	if  url[len(url)-1] != '/' {
		url = url + "/"
		//test()
		
	}
}