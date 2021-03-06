// Code generated by go-swagger; DO NOT EDIT.

package poll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/wunari/easypoll-backend/models"
)

// PatchPollByIDOKCode is the HTTP code returned for type PatchPollByIDOK
const PatchPollByIDOKCode int = 200

/*PatchPollByIDOK Poll updated

swagger:response patchPollByIdOK
*/
type PatchPollByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Poll `json:"body,omitempty"`
}

// NewPatchPollByIDOK creates PatchPollByIDOK with default headers values
func NewPatchPollByIDOK() *PatchPollByIDOK {

	return &PatchPollByIDOK{}
}

// WithPayload adds the payload to the patch poll by Id o k response
func (o *PatchPollByIDOK) WithPayload(payload *models.Poll) *PatchPollByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch poll by Id o k response
func (o *PatchPollByIDOK) SetPayload(payload *models.Poll) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPollByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPollByIDBadRequestCode is the HTTP code returned for type PatchPollByIDBadRequest
const PatchPollByIDBadRequestCode int = 400

/*PatchPollByIDBadRequest Invalid input

swagger:response patchPollByIdBadRequest
*/
type PatchPollByIDBadRequest struct {
}

// NewPatchPollByIDBadRequest creates PatchPollByIDBadRequest with default headers values
func NewPatchPollByIDBadRequest() *PatchPollByIDBadRequest {

	return &PatchPollByIDBadRequest{}
}

// WriteResponse to the client
func (o *PatchPollByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PatchPollByIDNotFoundCode is the HTTP code returned for type PatchPollByIDNotFound
const PatchPollByIDNotFoundCode int = 404

/*PatchPollByIDNotFound Resource not found

swagger:response patchPollByIdNotFound
*/
type PatchPollByIDNotFound struct {
}

// NewPatchPollByIDNotFound creates PatchPollByIDNotFound with default headers values
func NewPatchPollByIDNotFound() *PatchPollByIDNotFound {

	return &PatchPollByIDNotFound{}
}

// WriteResponse to the client
func (o *PatchPollByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
