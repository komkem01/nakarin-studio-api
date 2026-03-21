package prefix

import "errors"

var ErrPrefixNotFound = errors.New("prefix not found")
var ErrPrefixAlreadyExists = errors.New("prefix already exists")
var ErrPrefixInvalidInput = errors.New("invalid input for prefix")
var ErrPrefixDeleteFailed = errors.New("failed to delete prefix")
var ErrPrefixUpdateFailed = errors.New("failed to update prefix")
