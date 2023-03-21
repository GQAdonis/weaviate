//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/weaviate/weaviate/entities/models"
)

// WeaviateRootReader is a Reader for the WeaviateRoot structure.
type WeaviateRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWeaviateRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWeaviateRootOK creates a WeaviateRootOK with default headers values
func NewWeaviateRootOK() *WeaviateRootOK {
	return &WeaviateRootOK{}
}

/*
WeaviateRootOK describes a response with status code 200, with default header values.

Weaviate is alive and ready to serve content
*/
type WeaviateRootOK struct {
	Payload *WeaviateRootOKBody
}

// IsSuccess returns true when this weaviate root o k response has a 2xx status code
func (o *WeaviateRootOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this weaviate root o k response has a 3xx status code
func (o *WeaviateRootOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this weaviate root o k response has a 4xx status code
func (o *WeaviateRootOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this weaviate root o k response has a 5xx status code
func (o *WeaviateRootOK) IsServerError() bool {
	return false
}

// IsCode returns true when this weaviate root o k response a status code equal to that given
func (o *WeaviateRootOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the weaviate root o k response
func (o *WeaviateRootOK) Code() int {
	return 200
}

func (o *WeaviateRootOK) Error() string {
	return fmt.Sprintf("[GET /][%d] weaviateRootOK  %+v", 200, o.Payload)
}

func (o *WeaviateRootOK) String() string {
	return fmt.Sprintf("[GET /][%d] weaviateRootOK  %+v", 200, o.Payload)
}

func (o *WeaviateRootOK) GetPayload() *WeaviateRootOKBody {
	return o.Payload
}

func (o *WeaviateRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(WeaviateRootOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
WeaviateRootOKBody weaviate root o k body
swagger:model WeaviateRootOKBody
*/
type WeaviateRootOKBody struct {

	// links
	Links []*models.Link `json:"links"`
}

// Validate validates this weaviate root o k body
func (o *WeaviateRootOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateRootOKBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	for i := 0; i < len(o.Links); i++ {
		if swag.IsZero(o.Links[i]) { // not required
			continue
		}

		if o.Links[i] != nil {
			if err := o.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weaviateRootOK" + "." + "links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weaviateRootOK" + "." + "links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this weaviate root o k body based on the context it is used
func (o *WeaviateRootOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateRootOKBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Links); i++ {

		if o.Links[i] != nil {
			if err := o.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weaviateRootOK" + "." + "links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weaviateRootOK" + "." + "links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateRootOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateRootOKBody) UnmarshalBinary(b []byte) error {
	var res WeaviateRootOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
