package errors

import (
	"github.com/giantswarm/microerror"
)

var InvalidConfigError = &microerror.Error{
	Kind: "InvalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == InvalidConfigError
}

var NotFoundError = &microerror.Error{
	Kind: "NotFoundError",
}

// IsNotFound asserts notFoundError
func IsNotFound(err error) bool {
	return microerror.Cause(err) == NotFoundError
}

var BadRequestError = &microerror.Error{
	Kind: "BadRequestError",
}

// IsBadRequest asserts badRequestError.
func IsBadRequest(err error) bool {
	return microerror.Cause(err) == BadRequestError
}
