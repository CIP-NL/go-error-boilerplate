// Package go_error_boilerplate defines a boilerplate for the errors used in CIP go projects.
package go_error_boilerplate

import (
	"fmt"
)

// ErrorInterface should be a return parameter from a function instead of error.
type ErrorInterface interface {
	error
	IsNil() bool
	Code() string
	Kind() Kind
	Public() (string, bool)
	Private() ([]string, bool)
	Retry() bool
}

// Kind defines a certain class of error, and should remain constant per package to allow for error handling.
type Kind uint8

// An enumeration. Make sure the first one is None, all others may be changed depending on package.
const (
	None          Kind = iota
	Other              // Unclassified error. This value is not printed in the error message.
	Invalid            // Invalid operation for this type of item.
	Permission         // Permission denied.
	IO                 // External I/O error such as network failure.
	Exist              // Item already exists.
	NotExist           // Item does not exist.
	IsDir              // Item is a directory.
	NotDir             // Item is not a directory.
	NotEmpty           // Directory not empty.
	Private            // Information withheld.
	Internal           // Internal error or inconsistency.
	CannotDecrypt      // No wrapped key for user with read access.
	Transient          // A transient error.
	BrokenLink         // Link target does not exist.
)

// The error struct to be
type Error struct {
	code    string
	kind    Kind
	public  string
	retry   bool
	private []string
	error
}

// IsNil is used to verify if the error is nil or not, since a nil interface != nil
func (e *Error) IsNil() bool {
	if e.code != "" {
		return false
	}

	if e.kind != 0 {
		return false
	}
	return true
}

// Code() is the getter for code.
func (e *Error) Code() string {
	if e.code != "" {
		return e.code
	} else {
		panic("Programming error is nil!")
	}
}

// Kind() is the getter for kind.
func (e *Error) Kind() Kind {
	if e.kind != None {
		return e.kind
	}
	panic("Programming error is nil!")
}

// Public is the getter for public
func (e *Error) Public() (string, bool) {
	if e.public != "" {
		return e.public, true
	}
	return e.public, false
}

// Private() is the getter for private
func (e *Error) Private() ([]string, bool) {
	if len(e.private) > 0 {
		return e.private, true
	}
	return e.private, false
}

// Retry() is the getter for retry
func (e *Error) Retry() bool {
	return e.retry
}

func NewError(code string, kind Kind, public string, retry bool, private ...interface{}) ErrorInterface {

	privates := make([]string, len(private))
	for _, i := range private {
		privates = append(privates, fmt.Sprintf("%v", i))
	}

	return &Error{
		code:    code,
		kind:    kind,
		public:  public,
		retry:   retry,
		private: privates,
	}
}
