// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsCreateReader is a Reader for the ObjectsCreate structure.
type ObjectsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewObjectsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewObjectsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObjectsCreateOK creates a ObjectsCreateOK with default headers values
func NewObjectsCreateOK() *ObjectsCreateOK {
	return &ObjectsCreateOK{}
}

/*
ObjectsCreateOK describes a response with status code 200, with default header values.

Object created.
*/
type ObjectsCreateOK struct {
	Payload *models.Object
}

// IsSuccess returns true when this objects create o k response has a 2xx status code
func (o *ObjectsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this objects create o k response has a 3xx status code
func (o *ObjectsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects create o k response has a 4xx status code
func (o *ObjectsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects create o k response has a 5xx status code
func (o *ObjectsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this objects create o k response a status code equal to that given
func (o *ObjectsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the objects create o k response
func (o *ObjectsCreateOK) Code() int {
	return 200
}

func (o *ObjectsCreateOK) Error() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateOK  %+v", 200, o.Payload)
}

func (o *ObjectsCreateOK) String() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateOK  %+v", 200, o.Payload)
}

func (o *ObjectsCreateOK) GetPayload() *models.Object {
	return o.Payload
}

func (o *ObjectsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Object)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsCreateBadRequest creates a ObjectsCreateBadRequest with default headers values
func NewObjectsCreateBadRequest() *ObjectsCreateBadRequest {
	return &ObjectsCreateBadRequest{}
}

/*
ObjectsCreateBadRequest describes a response with status code 400, with default header values.

Malformed request.
*/
type ObjectsCreateBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects create bad request response has a 2xx status code
func (o *ObjectsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects create bad request response has a 3xx status code
func (o *ObjectsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects create bad request response has a 4xx status code
func (o *ObjectsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects create bad request response has a 5xx status code
func (o *ObjectsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this objects create bad request response a status code equal to that given
func (o *ObjectsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the objects create bad request response
func (o *ObjectsCreateBadRequest) Code() int {
	return 400
}

func (o *ObjectsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsCreateBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsCreateUnauthorized creates a ObjectsCreateUnauthorized with default headers values
func NewObjectsCreateUnauthorized() *ObjectsCreateUnauthorized {
	return &ObjectsCreateUnauthorized{}
}

/*
ObjectsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsCreateUnauthorized struct {
}

// IsSuccess returns true when this objects create unauthorized response has a 2xx status code
func (o *ObjectsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects create unauthorized response has a 3xx status code
func (o *ObjectsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects create unauthorized response has a 4xx status code
func (o *ObjectsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects create unauthorized response has a 5xx status code
func (o *ObjectsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this objects create unauthorized response a status code equal to that given
func (o *ObjectsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the objects create unauthorized response
func (o *ObjectsCreateUnauthorized) Code() int {
	return 401
}

func (o *ObjectsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateUnauthorized ", 401)
}

func (o *ObjectsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateUnauthorized ", 401)
}

func (o *ObjectsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsCreateForbidden creates a ObjectsCreateForbidden with default headers values
func NewObjectsCreateForbidden() *ObjectsCreateForbidden {
	return &ObjectsCreateForbidden{}
}

/*
ObjectsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ObjectsCreateForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects create forbidden response has a 2xx status code
func (o *ObjectsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects create forbidden response has a 3xx status code
func (o *ObjectsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects create forbidden response has a 4xx status code
func (o *ObjectsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects create forbidden response has a 5xx status code
func (o *ObjectsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this objects create forbidden response a status code equal to that given
func (o *ObjectsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the objects create forbidden response
func (o *ObjectsCreateForbidden) Code() int {
	return 403
}

func (o *ObjectsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsCreateUnprocessableEntity creates a ObjectsCreateUnprocessableEntity with default headers values
func NewObjectsCreateUnprocessableEntity() *ObjectsCreateUnprocessableEntity {
	return &ObjectsCreateUnprocessableEntity{}
}

/*
ObjectsCreateUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type ObjectsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects create unprocessable entity response has a 2xx status code
func (o *ObjectsCreateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects create unprocessable entity response has a 3xx status code
func (o *ObjectsCreateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects create unprocessable entity response has a 4xx status code
func (o *ObjectsCreateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects create unprocessable entity response has a 5xx status code
func (o *ObjectsCreateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this objects create unprocessable entity response a status code equal to that given
func (o *ObjectsCreateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the objects create unprocessable entity response
func (o *ObjectsCreateUnprocessableEntity) Code() int {
	return 422
}

func (o *ObjectsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsCreateUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsCreateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsCreateInternalServerError creates a ObjectsCreateInternalServerError with default headers values
func NewObjectsCreateInternalServerError() *ObjectsCreateInternalServerError {
	return &ObjectsCreateInternalServerError{}
}

/*
ObjectsCreateInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects create internal server error response has a 2xx status code
func (o *ObjectsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects create internal server error response has a 3xx status code
func (o *ObjectsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects create internal server error response has a 4xx status code
func (o *ObjectsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects create internal server error response has a 5xx status code
func (o *ObjectsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this objects create internal server error response a status code equal to that given
func (o *ObjectsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the objects create internal server error response
func (o *ObjectsCreateInternalServerError) Code() int {
	return 500
}

func (o *ObjectsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /objects][%d] objectsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
