// Package go_error_boilerplate defines a boilerplate for the errors used in CIP go projects.
package go_error_boilerplate

// Error should be a return parameter from a function instead of error.
type Error interface {
	error
	Code() string
	Kind() Kind
	Public() (string, bool)
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

// E is the the error struct to be returned by a NewError
type E struct {
	code    string
	kind    Kind
	public  string
	retry   bool
	error
}

// Code() is the getter for code. Code is used to identify errors and may be compared against.
func (e *E) Code() string {
	if e.code != "" {
		return e.code
	} else {
		panic("ProgrammingError: error is nil!")
	}
}

// Kind() is the getter for kind. Kind is the general class of the error and may be compared against.
func (e *E) Kind() Kind {
	if e.kind != None {
		return e.kind
	}
	panic("ProgrammingError: error is nil!")
}

// Public is the getter for public. It is the human readable interface of the error and should not be compared against.
func (e *E) Public() (string, bool) {
	if e.public != "" {
		return e.public, true
	}
	return e.public, false
}

func (e *E) Error() (string, bool) {
	return e.Public()
}

// Retry() is the getter for retry. It indicates to the calling function/party that they may retry with exponential backoff.
func (e *E) Retry() bool {
	return e.retry
}

// NewError returns a struct of interface Error.
func NewError(code string, kind Kind, public string, retry bool) Error {

	return &E{
		code:   code,
		kind:   kind,
		public: public,
		retry:  retry,
	}
}
