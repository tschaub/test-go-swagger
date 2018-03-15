// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/tschaub/test-go-swagger/gen/models"
)

// PostTestOKCode is the HTTP code returned for type PostTestOK
const PostTestOKCode int = 200

/*PostTestOK It works

swagger:response postTestOK
*/
type PostTestOK struct {

	/*
	  In: Body
	*/
	Payload *models.PostTestOKBody `json:"body,omitempty"`
}

// NewPostTestOK creates PostTestOK with default headers values
func NewPostTestOK() *PostTestOK {
	return &PostTestOK{}
}

// WithPayload adds the payload to the post test o k response
func (o *PostTestOK) WithPayload(payload *models.PostTestOKBody) *PostTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post test o k response
func (o *PostTestOK) SetPayload(payload *models.PostTestOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
