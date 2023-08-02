// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudCloudinstancesVolumesPostParams creates a new PcloudCloudinstancesVolumesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesVolumesPostParams() *PcloudCloudinstancesVolumesPostParams {
	return &PcloudCloudinstancesVolumesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesVolumesPostParamsWithTimeout creates a new PcloudCloudinstancesVolumesPostParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesVolumesPostParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesPostParams {
	return &PcloudCloudinstancesVolumesPostParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesVolumesPostParamsWithContext creates a new PcloudCloudinstancesVolumesPostParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesVolumesPostParamsWithContext(ctx context.Context) *PcloudCloudinstancesVolumesPostParams {
	return &PcloudCloudinstancesVolumesPostParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesVolumesPostParamsWithHTTPClient creates a new PcloudCloudinstancesVolumesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesVolumesPostParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesPostParams {
	return &PcloudCloudinstancesVolumesPostParams{
		HTTPClient: client,
	}
}

/*
PcloudCloudinstancesVolumesPostParams contains all the parameters to send to the API endpoint

	for the pcloud cloudinstances volumes post operation.

	Typically these are written to a http.Request.
*/
type PcloudCloudinstancesVolumesPostParams struct {

	/* Body.

	   Parameters for the creation of a new data volume
	*/
	Body *models.CreateDataVolume

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances volumes post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesVolumesPostParams) WithDefaults() *PcloudCloudinstancesVolumesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances volumes post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesVolumesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) WithContext(ctx context.Context) *PcloudCloudinstancesVolumesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) WithBody(body *models.CreateDataVolume) *PcloudCloudinstancesVolumesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) SetBody(body *models.CreateDataVolume) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesVolumesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances volumes post params
func (o *PcloudCloudinstancesVolumesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesVolumesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
