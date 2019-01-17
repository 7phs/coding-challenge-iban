package helper

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrList(t *testing.T) {
	assert.NoError(t, ErrList{}.Result(), "no error")

	errList := ErrList{}
	errList.Add(nil)

	assert.NoError(t, errList.Result(), "no error")

	errList.Add(os.ErrInvalid)
	errList.Add(os.ErrExist)

	assert.Error(t, errList.Result(), "error")
	assert.Equal(t,
		fmt.Sprintf("%v; %v", os.ErrInvalid, os.ErrExist),
		errList.Error())
}
