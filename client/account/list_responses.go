// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/wecom/models"
)

// ListReader is a Reader for the List structure.
type ListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOK creates a ListOK with default headers values
func NewListOK() *ListOK {
	return &ListOK{}
}

/*
ListOK describes a response with status code 200, with default header values.

list customer service account response
*/
type ListOK struct {
	Payload *models.ListCustomerServiceAccountReponse
}

// IsSuccess returns true when this list o k response has a 2xx status code
func (o *ListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list o k response has a 3xx status code
func (o *ListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list o k response has a 4xx status code
func (o *ListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list o k response has a 5xx status code
func (o *ListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list o k response a status code equal to that given
func (o *ListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list o k response
func (o *ListOK) Code() int {
	return 200
}

func (o *ListOK) Error() string {
	return fmt.Sprintf("[POST /kf/account/list][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) String() string {
	return fmt.Sprintf("[POST /kf/account/list][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) GetPayload() *models.ListCustomerServiceAccountReponse {
	return o.Payload
}

func (o *ListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCustomerServiceAccountReponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDefault creates a ListDefault with default headers values
func NewListDefault(code int) *ListDefault {
	return &ListDefault{
		_statusCode: code,
	}
}

/*
ListDefault describes a response with status code -1, with default header values.

unexpected error
*/
type ListDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this list default response has a 2xx status code
func (o *ListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list default response has a 3xx status code
func (o *ListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list default response has a 4xx status code
func (o *ListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list default response has a 5xx status code
func (o *ListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list default response a status code equal to that given
func (o *ListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list default response
func (o *ListDefault) Code() int {
	return o._statusCode
}

func (o *ListDefault) Error() string {
	return fmt.Sprintf("[POST /kf/account/list][%d] list default  %+v", o._statusCode, o.Payload)
}

func (o *ListDefault) String() string {
	return fmt.Sprintf("[POST /kf/account/list][%d] list default  %+v", o._statusCode, o.Payload)
}

func (o *ListDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}