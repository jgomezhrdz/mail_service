package custom_errors

import "errors"

var ErrNotFound = errors.New("resource was not found")
