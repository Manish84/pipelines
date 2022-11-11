// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/pipeline_model"
)

// GetPipelineV1Reader is a Reader for the GetPipelineV1 structure.
type GetPipelineV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPipelineV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPipelineV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelineV1OK creates a GetPipelineV1OK with default headers values
func NewGetPipelineV1OK() *GetPipelineV1OK {
	return &GetPipelineV1OK{}
}

/*GetPipelineV1OK handles this case with default header values.

A successful response.
*/
type GetPipelineV1OK struct {
	Payload *pipeline_model.V1beta1Pipeline
}

func (o *GetPipelineV1OK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/pipelines/{id}][%d] getPipelineV1OK  %+v", 200, o.Payload)
}

func (o *GetPipelineV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V1beta1Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineV1Default creates a GetPipelineV1Default with default headers values
func NewGetPipelineV1Default(code int) *GetPipelineV1Default {
	return &GetPipelineV1Default{
		_statusCode: code,
	}
}

/*GetPipelineV1Default handles this case with default header values.

GetPipelineV1Default get pipeline v1 default
*/
type GetPipelineV1Default struct {
	_statusCode int

	Payload *pipeline_model.V1beta1Status
}

// Code gets the status code for the get pipeline v1 default response
func (o *GetPipelineV1Default) Code() int {
	return o._statusCode
}

func (o *GetPipelineV1Default) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/pipelines/{id}][%d] GetPipelineV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetPipelineV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V1beta1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
