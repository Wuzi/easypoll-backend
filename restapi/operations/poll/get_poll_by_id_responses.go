// Code generated by go-swagger; DO NOT EDIT.

package poll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/wunari/easypoll-backend/models"
)

// GetPollByIDOKCode is the HTTP code returned for type GetPollByIDOK
const GetPollByIDOKCode int = 200

/*GetPollByIDOK Poll fetched successfully

swagger:response getPollByIdOK
*/
type GetPollByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Poll `json:"body,omitempty"`
}

// NewGetPollByIDOK creates GetPollByIDOK with default headers values
func NewGetPollByIDOK() *GetPollByIDOK {

	return &GetPollByIDOK{}
}

// WithPayload adds the payload to the get poll by Id o k response
func (o *GetPollByIDOK) WithPayload(payload *models.Poll) *GetPollByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get poll by Id o k response
func (o *GetPollByIDOK) SetPayload(payload *models.Poll) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPollByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPollByIDNotFoundCode is the HTTP code returned for type GetPollByIDNotFound
const GetPollByIDNotFoundCode int = 404

/*GetPollByIDNotFound Resource not found

swagger:response getPollByIdNotFound
*/
type GetPollByIDNotFound struct {
}

// NewGetPollByIDNotFound creates GetPollByIDNotFound with default headers values
func NewGetPollByIDNotFound() *GetPollByIDNotFound {

	return &GetPollByIDNotFound{}
}

// WriteResponse to the client
func (o *GetPollByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
