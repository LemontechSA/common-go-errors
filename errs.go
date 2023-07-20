package errs

import (
	"errors"
	"fmt"
	"net/http"
)

// Errors wrapper structure
type ErrorWrapper struct {
	Action  string            `json:"action"`  // Human readable action that origin the error.
	Message string            `json:"message"` // Human readable message for clients.
	Payload map[string]string `json:"payload"` // Extra information for logs purposes.
	Code    int               `json:"-"`       // HTTP Status code. `-` is used to skip json marshaling.
	Err     error             `json:"-"`       // The original error. Same reason as above.
}

// Returns Message if Err is nil.
func (err ErrorWrapper) Error() string {
	// guard against panics
	if err.Err != nil {
		return err.Err.Error()
	}

	return err.Message
}

// Implements the errors.Unwrap interface
func (err ErrorWrapper) Unwrap() error {
	return err.Err // Returns inner error
}

// Returns the inner most ErrorWrapper
func (err ErrorWrapper) Dig() ErrorWrapper {
	var ew ErrorWrapper

	if errors.As(err.Err, &ew) {
		// Recursively digs until wrapper error is not in which case it will stop
		return ew.Dig()
	}

	return err
}

// Add a value to payload - if the key already exists the value will be overrite
func (err *ErrorWrapper) AddPayloadValue(key string, value string) {
	if err.Payload == nil {
		err.Payload = map[string]string{key: value}
	} else {
		err.Payload[key] = value
	}
}

// Adds values to payload - if a key already exists the value will be overrite
func (err *ErrorWrapper) AddPayloadValues(values map[string]string) {
	if err.Payload == nil {
		err.Payload = values
	} else {
		for k, v := range values {
			err.Payload[k] = v
		}
	}
}

// Returns the values of action and message as json
func (err ErrorWrapper) AsJSONResponse() map[string]string {
	return map[string]string{
		"action":  err.Action,
		"message": err.Message,
	}
}

// Returns the inner ErrorWrapper or a generic one
func DecodeError(err error) ErrorWrapper {
	var ew ErrorWrapper

	if errors.As(err, &ew) {
		return ew
	}

	return ErrorWrapper{
		Action:  "generic",
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
		Err:     err,
		Payload: nil,
	}
}

// Returns a generic wrapped error
func NewErrorWrapper(
	action string,
	message string,
	code int,
	err error,
	payload map[string]string,
) error {
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Bad Request - Http code 400
func NewBadRequestError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusBadRequest
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Unauthorized - Http code 401
func NewUnauthorizedError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusUnauthorized
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Payment Required - Http code 402
func NewPaymentRequiredError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusPaymentRequired
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Forbidden - Http code 403
func NewForbiddenError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusForbidden
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Not Found - Http code 404
func NewNotFoundError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusNotFound
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Unprocessable Entity - Http code 422
func NewUnprocessableEntityError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusUnprocessableEntity
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Internal Server Error - Http code 500
func NewInternalServerError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusInternalServerError
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Not Implemented - Http code 501
func NewNotImplementedError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusNotImplemented
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Bad Gateway - Http code 502
func NewBadGatewayError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusBadGateway
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Service Unavailable - Http code 503
func NewServiceUnavailableError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusServiceUnavailable
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}

// Returns a wrapped error of type Gateway Timeout - Http code 504
func NewGatewayTimeoutError(
	action string,
	message string,
	err error,
	payload map[string]string,
) error {
	code := http.StatusGatewayTimeout
	ew := ErrorWrapper{
		Action:  action,
		Message: message,
		Code:    code,
		Err:     err,
		Payload: payload,
	}

	ew.AddPayloadValues(map[string]string{
		"code":          fmt.Sprint(code),
		"human_message": message,
	})

	return ew
}
