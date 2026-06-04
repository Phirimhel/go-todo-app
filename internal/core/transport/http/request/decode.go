package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New(validator.WithRequiredStructEnabled())

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {

	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		return fmt.Errorf("decode json: %v, %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	var err error

	// check if DTO has its own rules of validation
	v, ok := dest.(validatable)
	if ok {
		err = v.Validate()
	} else {
		// if it hasn't: use stardat validation method
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("[DTO validation]: request validation: %v, %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
