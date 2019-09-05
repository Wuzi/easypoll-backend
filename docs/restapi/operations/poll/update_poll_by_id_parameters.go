// Code generated by go-swagger; DO NOT EDIT.

package poll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/wunari/easypoll-backend/docs/models"
)

// NewUpdatePollByIDParams creates a new UpdatePollByIDParams object
// no default values defined in spec.
func NewUpdatePollByIDParams() UpdatePollByIDParams {

	return UpdatePollByIDParams{}
}

// UpdatePollByIDParams contains all the bound params for the update poll by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePollById
type UpdatePollByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*New poll data
	  Required: true
	  In: body
	*/
	Body *models.Poll
	/*The id of the poll
	  Required: true
	  In: path
	*/
	ID float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdatePollByIDParams() beforehand.
func (o *UpdatePollByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Poll
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdatePollByIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "float64", raw)
	}
	o.ID = value

	return nil
}
