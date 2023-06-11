// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/wecom/models"
)

// SendReader is a Reader for the Send structure.
type SendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSendOK creates a SendOK with default headers values
func NewSendOK() *SendOK {
	return &SendOK{}
}

/*
SendOK describes a response with status code 200, with default header values.

send message response
*/
type SendOK struct {
	Payload *models.SendReponse
}

// IsSuccess returns true when this send o k response has a 2xx status code
func (o *SendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send o k response has a 3xx status code
func (o *SendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send o k response has a 4xx status code
func (o *SendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this send o k response has a 5xx status code
func (o *SendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this send o k response a status code equal to that given
func (o *SendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the send o k response
func (o *SendOK) Code() int {
	return 200
}

func (o *SendOK) Error() string {
	return fmt.Sprintf("[POST /kf/send_msg][%d] sendOK  %+v", 200, o.Payload)
}

func (o *SendOK) String() string {
	return fmt.Sprintf("[POST /kf/send_msg][%d] sendOK  %+v", 200, o.Payload)
}

func (o *SendOK) GetPayload() *models.SendReponse {
	return o.Payload
}

func (o *SendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SendReponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendDefault creates a SendDefault with default headers values
func NewSendDefault(code int) *SendDefault {
	return &SendDefault{
		_statusCode: code,
	}
}

/*
SendDefault describes a response with status code -1, with default header values.

unexpected error
*/
type SendDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this send default response has a 2xx status code
func (o *SendDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this send default response has a 3xx status code
func (o *SendDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this send default response has a 4xx status code
func (o *SendDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this send default response has a 5xx status code
func (o *SendDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this send default response a status code equal to that given
func (o *SendDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the send default response
func (o *SendDefault) Code() int {
	return o._statusCode
}

func (o *SendDefault) Error() string {
	return fmt.Sprintf("[POST /kf/send_msg][%d] send default  %+v", o._statusCode, o.Payload)
}

func (o *SendDefault) String() string {
	return fmt.Sprintf("[POST /kf/send_msg][%d] send default  %+v", o._statusCode, o.Payload)
}

func (o *SendDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *SendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}