package utils

import "errors"

// Tipuri de erori
var ErrNoDocument = errors.New("mongo: no documents in result") // "mongo: no documents in result"
var ErrInvalidID = errors.New("the provided hex string is not a valid ObjectID")
var ErrUnexpect = errors.New("unexpected error")
var ErrHashFailed = errors.New("password hashing failed")
var ErrWrongPassword = errors.New("wrong password")
var ErrDuplicateDocument = errors.New("duplicate document")
