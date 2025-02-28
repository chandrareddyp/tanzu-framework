// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ExportTKGConfigForVsphereOKCode is the HTTP code returned for type ExportTKGConfigForVsphereOK
const ExportTKGConfigForVsphereOKCode int = 200

/*ExportTKGConfigForVsphereOK Generated TKG configuration successfully

swagger:response exportTKGConfigForVsphereOK
*/
type ExportTKGConfigForVsphereOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewExportTKGConfigForVsphereOK creates ExportTKGConfigForVsphereOK with default headers values
func NewExportTKGConfigForVsphereOK() *ExportTKGConfigForVsphereOK {

	return &ExportTKGConfigForVsphereOK{}
}

// WithPayload adds the payload to the export t k g config for vsphere o k response
func (o *ExportTKGConfigForVsphereOK) WithPayload(payload string) *ExportTKGConfigForVsphereOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export t k g config for vsphere o k response
func (o *ExportTKGConfigForVsphereOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportTKGConfigForVsphereOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ExportTKGConfigForVsphereBadRequestCode is the HTTP code returned for type ExportTKGConfigForVsphereBadRequest
const ExportTKGConfigForVsphereBadRequestCode int = 400

/*ExportTKGConfigForVsphereBadRequest Bad request

swagger:response exportTKGConfigForVsphereBadRequest
*/
type ExportTKGConfigForVsphereBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewExportTKGConfigForVsphereBadRequest creates ExportTKGConfigForVsphereBadRequest with default headers values
func NewExportTKGConfigForVsphereBadRequest() *ExportTKGConfigForVsphereBadRequest {

	return &ExportTKGConfigForVsphereBadRequest{}
}

// WithPayload adds the payload to the export t k g config for vsphere bad request response
func (o *ExportTKGConfigForVsphereBadRequest) WithPayload(payload *models.Error) *ExportTKGConfigForVsphereBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export t k g config for vsphere bad request response
func (o *ExportTKGConfigForVsphereBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportTKGConfigForVsphereBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ExportTKGConfigForVsphereUnauthorizedCode is the HTTP code returned for type ExportTKGConfigForVsphereUnauthorized
const ExportTKGConfigForVsphereUnauthorizedCode int = 401

/*ExportTKGConfigForVsphereUnauthorized Incorrect credentials

swagger:response exportTKGConfigForVsphereUnauthorized
*/
type ExportTKGConfigForVsphereUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewExportTKGConfigForVsphereUnauthorized creates ExportTKGConfigForVsphereUnauthorized with default headers values
func NewExportTKGConfigForVsphereUnauthorized() *ExportTKGConfigForVsphereUnauthorized {

	return &ExportTKGConfigForVsphereUnauthorized{}
}

// WithPayload adds the payload to the export t k g config for vsphere unauthorized response
func (o *ExportTKGConfigForVsphereUnauthorized) WithPayload(payload *models.Error) *ExportTKGConfigForVsphereUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export t k g config for vsphere unauthorized response
func (o *ExportTKGConfigForVsphereUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportTKGConfigForVsphereUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ExportTKGConfigForVsphereInternalServerErrorCode is the HTTP code returned for type ExportTKGConfigForVsphereInternalServerError
const ExportTKGConfigForVsphereInternalServerErrorCode int = 500

/*ExportTKGConfigForVsphereInternalServerError Internal server error

swagger:response exportTKGConfigForVsphereInternalServerError
*/
type ExportTKGConfigForVsphereInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewExportTKGConfigForVsphereInternalServerError creates ExportTKGConfigForVsphereInternalServerError with default headers values
func NewExportTKGConfigForVsphereInternalServerError() *ExportTKGConfigForVsphereInternalServerError {

	return &ExportTKGConfigForVsphereInternalServerError{}
}

// WithPayload adds the payload to the export t k g config for vsphere internal server error response
func (o *ExportTKGConfigForVsphereInternalServerError) WithPayload(payload *models.Error) *ExportTKGConfigForVsphereInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export t k g config for vsphere internal server error response
func (o *ExportTKGConfigForVsphereInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportTKGConfigForVsphereInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
