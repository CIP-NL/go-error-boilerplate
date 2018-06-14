package go_error_boilerplate

// Boilerplate error code for use in any go project in the CIP project. Errors should be generated as follows:
// a function returns an ErrorInterface, which is checked for IsNil(). Don't directly check by accessing code, kind or error,
// as a nil interface != nil.
//
//--REQUIRED --
// Code() is the getter for code, and should return a description that does not change and can
// be used to specifically catch an error.
//
//--REQUIRED --
// Kind() is the getter for kind, a general classification of the Error and represented by an enumeration. 0 means the kind is None.
//
//--REQUIRED --
// Error() is the getter for the string accessor, which should return a human readable representation of the error.
//
//--NOT REQUIRED --
// Public() is the getter for the public information on the error, which could be boiled up to the end user if needed.
// this means that for security reasons, public may not contain any sensitive information. public may be set to a nil string
// if the error is not applicable to end users.

type ErrorInterface interface {
	error
	IsNil() bool
	Code() string
	Kind() Kind
	Public() (string, bool)
}

type Kind uint8

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

type Error struct {
	code   string
	kind   Kind
	public string
	error
}

func (e *Error) IsNil() bool {
	if e.code != "" {
		return false
	}

	if e.kind != 0 {
		return false
	}
	return true
}

func (e *Error) Code() string {
	if e.code != "" {
		return e.code
	} else {
		panic("Programming error is nil!")
	}
}

func (e *Error) Kind() Kind {
	if e.kind != None {
		return e.kind
	}
	panic("Programming error is nil!")
}

func (e *Error) Public() (string, bool) {
	if e.public != "" {
		return e.public, true
	}
	return e.public, false
}
