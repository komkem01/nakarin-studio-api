package gender

import "errors"

var ErrGenderNotFound = errors.New("gender not found")
var ErrGenderAlreadyExists = errors.New("gender already exists")
var ErrGenderInvalidInput = errors.New("invalid input for gender")
var ErrGenderDeleteFailed = errors.New("failed to delete gender")
var ErrGenderUpdateFailed = errors.New("failed to update gender")
