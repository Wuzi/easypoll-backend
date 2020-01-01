// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/wunari/easypoll-backend/models"
)

// GetUserHandlerFunc turns a function with the right signature into a get user handler
type GetUserHandlerFunc func(GetUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserHandlerFunc) Handle(params GetUserParams) middleware.Responder {
	return fn(params)
}

// GetUserHandler interface for that can handle valid get user params
type GetUserHandler interface {
	Handle(GetUserParams) middleware.Responder
}

// NewGetUser creates a new http.Handler for the get user operation
func NewGetUser(ctx *middleware.Context, handler GetUserHandler) *GetUser {
	return &GetUser{Context: ctx, Handler: handler}
}

/*GetUser swagger:route GET /user auth getUser

Get the logged user info

Get the authenticated user account details

*/
type GetUser struct {
	Context *middleware.Context
	Handler GetUserHandler
}

func (o *GetUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUserOKBody get user o k body
// swagger:model GetUserOKBody
type GetUserOKBody struct {
	models.User
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetUserOKBody) UnmarshalJSON(raw []byte) error {
	// GetUserOKBodyAO0
	var getUserOKBodyAO0 models.User
	if err := swag.ReadJSON(raw, &getUserOKBodyAO0); err != nil {
		return err
	}
	o.User = getUserOKBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetUserOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getUserOKBodyAO0, err := swag.WriteJSON(o.User)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getUserOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get user o k body
func (o *GetUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.User
	if err := o.User.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}