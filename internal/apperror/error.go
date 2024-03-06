package apperror

import "errors"

// InvalidUserErr user related error
var InvalidUserErr = errors.New("invalid user")
var MissingCartIdErr = errors.New("missing cartId in request")

// ServerErr is server end error
var ServerErr = errors.New("error in server end")
