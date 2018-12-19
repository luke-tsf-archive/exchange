package helpers

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v8"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	// convert error into validation error
	// invoke this method when it hit vadidation error
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)

		// Store error for each param
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// mapping error into an interface of key, value
func NewError(key string, err error) CommonError {
	commonError := CommonError{}
	commonError.Errors = make(map[string]interface{})
	commonError.Errors[key] = err
	return commonError
}
