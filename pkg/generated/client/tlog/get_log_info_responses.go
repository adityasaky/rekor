// Code generated by go-swagger; DO NOT EDIT.

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/projectrekor/rekor/pkg/generated/models"
)

// GetLogInfoReader is a Reader for the GetLogInfo structure.
type GetLogInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLogInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogInfoOK creates a GetLogInfoOK with default headers values
func NewGetLogInfoOK() *GetLogInfoOK {
	return &GetLogInfoOK{}
}

/*GetLogInfoOK handles this case with default header values.

A JSON object with the root hash and tree size as properties
*/
type GetLogInfoOK struct {
	Payload *models.LogInfo
}

func (o *GetLogInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/log][%d] getLogInfoOK  %+v", 200, o.Payload)
}

func (o *GetLogInfoOK) GetPayload() *models.LogInfo {
	return o.Payload
}

func (o *GetLogInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogInfoDefault creates a GetLogInfoDefault with default headers values
func NewGetLogInfoDefault(code int) *GetLogInfoDefault {
	return &GetLogInfoDefault{
		_statusCode: code,
	}
}

/*GetLogInfoDefault handles this case with default header values.

There was an internal error in the server while processing the request
*/
type GetLogInfoDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get log info default response
func (o *GetLogInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetLogInfoDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/log][%d] getLogInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogInfoDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}