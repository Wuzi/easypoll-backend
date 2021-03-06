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
	validate "github.com/go-openapi/validate"

	models "github.com/wunari/easypoll-backend/models"
)

// PostLoginHandlerFunc turns a function with the right signature into a post login handler
type PostLoginHandlerFunc func(PostLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostLoginHandlerFunc) Handle(params PostLoginParams) middleware.Responder {
	return fn(params)
}

// PostLoginHandler interface for that can handle valid post login params
type PostLoginHandler interface {
	Handle(PostLoginParams) middleware.Responder
}

// NewPostLogin creates a new http.Handler for the post login operation
func NewPostLogin(ctx *middleware.Context, handler PostLoginHandler) *PostLogin {
	return &PostLogin{Context: ctx, Handler: handler}
}

/*PostLogin swagger:route POST /login auth postLogin

Login to the application

Authenticates an user and returns a token to be used in requests

*/
type PostLogin struct {
	Context *middleware.Context
	Handler PostLoginHandler
}

func (o *PostLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostLoginParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostLoginBody post login body
// swagger:model PostLoginBody
type PostLoginBody struct {

	// Email of the account
	// Required: true
	Email *string `json:"email"`

	// Password of the account
	// Required: true
	// Format: password
	Password *strfmt.Password `json:"password"`
}

// Validate validates this post login body
func (o *PostLoginBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLoginBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *PostLoginBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginBody) UnmarshalBinary(b []byte) error {
	var res PostLoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostLoginOKBody post login o k body
// swagger:model PostLoginOKBody
type PostLoginOKBody struct {
	models.Token

	// user
	User *models.User `json:"user,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostLoginOKBody) UnmarshalJSON(raw []byte) error {
	// PostLoginOKBodyAO0
	var postLoginOKBodyAO0 models.Token
	if err := swag.ReadJSON(raw, &postLoginOKBodyAO0); err != nil {
		return err
	}
	o.Token = postLoginOKBodyAO0

	// PostLoginOKBodyAO1
	var dataPostLoginOKBodyAO1 struct {
		User *models.User `json:"user,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostLoginOKBodyAO1); err != nil {
		return err
	}

	o.User = dataPostLoginOKBodyAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostLoginOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postLoginOKBodyAO0, err := swag.WriteJSON(o.Token)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postLoginOKBodyAO0)

	var dataPostLoginOKBodyAO1 struct {
		User *models.User `json:"user,omitempty"`
	}

	dataPostLoginOKBodyAO1.User = o.User

	jsonDataPostLoginOKBodyAO1, errPostLoginOKBodyAO1 := swag.WriteJSON(dataPostLoginOKBodyAO1)
	if errPostLoginOKBodyAO1 != nil {
		return nil, errPostLoginOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostLoginOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post login o k body
func (o *PostLoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Token
	if err := o.Token.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLoginOKBody) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postLoginOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginOKBody) UnmarshalBinary(b []byte) error {
	var res PostLoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
