// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/run_model"
)

// RetryRunV1Reader is a Reader for the RetryRunV1 structure.
type RetryRunV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryRunV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetryRunV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRetryRunV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRetryRunV1OK creates a RetryRunV1OK with default headers values
func NewRetryRunV1OK() *RetryRunV1OK {
	return &RetryRunV1OK{}
}

/*RetryRunV1OK handles this case with default header values.

A successful response.
*/
type RetryRunV1OK struct {
	Payload interface{}
}

func (o *RetryRunV1OK) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/runs/{run_id}/retry][%d] retryRunV1OK  %+v", 200, o.Payload)
}

func (o *RetryRunV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryRunV1Default creates a RetryRunV1Default with default headers values
func NewRetryRunV1Default(code int) *RetryRunV1Default {
	return &RetryRunV1Default{
		_statusCode: code,
	}
}

/*RetryRunV1Default handles this case with default header values.

RetryRunV1Default retry run v1 default
*/
type RetryRunV1Default struct {
	_statusCode int

	Payload *run_model.V1beta1Status
}

// Code gets the status code for the retry run v1 default response
func (o *RetryRunV1Default) Code() int {
	return o._statusCode
}

func (o *RetryRunV1Default) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/runs/{run_id}/retry][%d] RetryRunV1 default  %+v", o._statusCode, o.Payload)
}

func (o *RetryRunV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V1beta1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
