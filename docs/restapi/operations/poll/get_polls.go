// Code generated by go-swagger; DO NOT EDIT.

package poll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPollsHandlerFunc turns a function with the right signature into a get polls handler
type GetPollsHandlerFunc func(GetPollsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPollsHandlerFunc) Handle(params GetPollsParams) middleware.Responder {
	return fn(params)
}

// GetPollsHandler interface for that can handle valid get polls params
type GetPollsHandler interface {
	Handle(GetPollsParams) middleware.Responder
}

// NewGetPolls creates a new http.Handler for the get polls operation
func NewGetPolls(ctx *middleware.Context, handler GetPollsHandler) *GetPolls {
	return &GetPolls{Context: ctx, Handler: handler}
}

/*GetPolls swagger:route GET /polls poll getPolls

Gets an array with all the polls

*/
type GetPolls struct {
	Context *middleware.Context
	Handler GetPollsHandler
}

func (o *GetPolls) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPollsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}