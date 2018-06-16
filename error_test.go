package go_error_boilerplate

import (
	"testing"
	"github.com/stretchr/testify/assert"

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
	"Yes",
	true,
)

func TestError_Code(t *testing.T) {
	assert.Equal(t,"test:1", err1.Code(),)
	assert.Equal(t,"test:2", err2.Code(),)
}

func TestError_Kind(t *testing.T) {
	assert.Equal(t, err1.Kind(), Other)
	assert.Equal(t, err2.Kind(), Private)
}


func TestError_Public(t *testing.T) {
	_, public := err1.Public()

	assert.False(t, public)

	_, public2 := err2.Public()
	assert.True(t, public2)

}

func TestError_Retry(t *testing.T) {
	retry := err1.Retry()
	assert.False(t, retry)

	retry2 := err2.Retry()
	assert.True(t, retry2)
}

func TestNewError(t *testing.T) {
	err := NewError(
		"test:1",
		Other,
		"",
		false,
	)

	assert.Error(t, err)
	assert.Equal(t, err.Kind(), Other)
}

// Checks if, when no error is returned, IsNil still works.
func TestError_IsNil2(t *testing.T) {
	err := func() Error {return nil}
	assert.Nil(t, err())
}