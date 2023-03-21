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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewObjectsHeadParams creates a new ObjectsHeadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewObjectsHeadParams() *ObjectsHeadParams {
	return &ObjectsHeadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsHeadParamsWithTimeout creates a new ObjectsHeadParams object
// with the ability to set a timeout on a request.
func NewObjectsHeadParamsWithTimeout(timeout time.Duration) *ObjectsHeadParams {
	return &ObjectsHeadParams{
		timeout: timeout,
	}
}

// NewObjectsHeadParamsWithContext creates a new ObjectsHeadParams object
// with the ability to set a context for a request.
func NewObjectsHeadParamsWithContext(ctx context.Context) *ObjectsHeadParams {
	return &ObjectsHeadParams{
		Context: ctx,
	}
}

// NewObjectsHeadParamsWithHTTPClient creates a new ObjectsHeadParams object
// with the ability to set a custom HTTPClient for a request.
func NewObjectsHeadParamsWithHTTPClient(client *http.Client) *ObjectsHeadParams {
	return &ObjectsHeadParams{
		HTTPClient: client,
	}
}

/*
ObjectsHeadParams contains all the parameters to send to the API endpoint

	for the objects head operation.

	Typically these are written to a http.Request.
*/
type ObjectsHeadParams struct {

	/* ID.

	   Unique ID of the Object.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the objects head params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectsHeadParams) WithDefaults() *ObjectsHeadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the objects head params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectsHeadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the objects head params
func (o *ObjectsHeadParams) WithTimeout(timeout time.Duration) *ObjectsHeadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects head params
func (o *ObjectsHeadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects head params
func (o *ObjectsHeadParams) WithContext(ctx context.Context) *ObjectsHeadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects head params
func (o *ObjectsHeadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects head params
func (o *ObjectsHeadParams) WithHTTPClient(client *http.Client) *ObjectsHeadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects head params
func (o *ObjectsHeadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the objects head params
func (o *ObjectsHeadParams) WithID(id strfmt.UUID) *ObjectsHeadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the objects head params
func (o *ObjectsHeadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsHeadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
