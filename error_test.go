package go_error_boilerplate

import (
	"testing"
)

var err1 = NewError(
	"test:1",
	Other,
	"",
	false,
)

var err2 = NewError(
	"test:2",
	Private,
	"",
	true,
	"accesstoken", "privatekey",
)

func TestError_Code(t *testing.T) {
	if err1.Code() != "test:1" {
		t.Error("Code 1 did not match")
	}

	if err2.Code() != "test:2" {
		t.Error("Code 1 did not match")
	}
}

func TestError_IsNil(t *testing.T) {
	if err1.IsNil() {
		t.Error("err1 is nil!")
	}

	if err2.IsNil() {
		t.Error("err2 is nil!")
	}
}

func TestError_Kind(t *testing.T) {
	if err1.Kind() != Other {
		t.Error("err1.Kind() should be Other!")
	}

	if err2.Kind() != Private {
		t.Error("err2.Kind() should be Private!")
	}
}

func TestError_Private(t *testing.T) {
	_, private := err1.Private()
	if private {
		t.Error("err1 is not private!")
	}

	_, private = err2.Private()
	if !private {
		t.Error("err2 is private!")
	}
}

func TestError_Public(t *testing.T) {
	_, public := err1.Public()
	if public {
		t.Error("err1 does not have public strings")
	}

	_, public = err2.Public()
	if public {
		t.Error("err2 does not have public strings")
	}
}

func TestError_Retry(t *testing.T) {
	retry := err1.Retry()
	if retry {
		t.Error("retry should be false.")
	}

	retry = err2.Retry()
	if !retry {
		t.Error("retry should be true.")
	}
}

func TestNewError(t *testing.T) {
	err := NewError(
		"test:1",
		Other,
		"",
		false,
	)

	if err.Kind() != Other {
		t.Error("NewError generated a nil error")
	}
}
