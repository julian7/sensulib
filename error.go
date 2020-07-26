package sensulib

import (
	"errors"
	"fmt"
)

const (
	OK = iota
	WARN
	CRIT
	UNKNOWN
)

// Error is a sensu-aware error message, which knows
// about its criticality
type Error struct {
	criticality int
	err         error
}

// Error returns a string representing the error and its
// criticality
func (serr *Error) Error() string {
	return fmt.Sprintf("%s: %v", serr.critString(), serr.err)
}

// Exit terminates run by returning the error
func (serr *Error) Exit() {
	fmt.Printf("%s\n", serr.Error())
	Exit(serr.criticality)
}

func args2err(args []interface{}) error {
	switch len(args) {
	case 0:
		return errors.New("(no message)")
	case 1:
		if errarg, ok := args[0].(error); ok {
			return errarg
		} else if errarg, ok := args[0].(string); ok {
			return errors.New(errarg)
		} else {
			return fmt.Errorf("%v", args[0])
		}
	}
	return fmt.Errorf(args[0].(string), args[1:]...)
}

// NewError creates a new Error
func NewError(crit int, args ...interface{}) *Error {
	err := args2err(args)
	return &Error{criticality: crit, err: err}
}

// Warn creates a Warning-level error
func Warn(args ...interface{}) *Error {
	return NewError(WARN, args...)
}

// Crit creates a Critical-level error
func Crit(args ...interface{}) *Error {
	return NewError(CRIT, args...)
}

// Ok creates a OK-level error
func Ok(args ...interface{}) *Error {
	return NewError(OK, args...)
}

// Unknown creates a Unknown-level error
func Unknown(args ...interface{}) *Error {
	return NewError(UNKNOWN, args...)
}

func (serr *Error) critString() string {
	switch serr.criticality {
	case OK:
		return "OK"
	case WARN:
		return "WARNING"
	case CRIT:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}
