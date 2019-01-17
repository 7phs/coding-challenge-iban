package helper

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/7phs/coding-challenge-iban/model/errCode"

	"github.com/stretchr/testify/assert"
)

func TestErrorRecordList_Error(t *testing.T) {
	var (
		exist     RespErrorRecordList
		existResp RespError
		expected  = RespErrorRecordList{
			{
				Id:   1,
				Desc: "err-true",
			},
			{
				Id:   2,
				Desc: "err-15",
			},
			{
				Id:   3,
				Desc: "3err-1",
			},
			{
				Id:   3,
				Desc: "3err-2",
			},
			{
				Id:   4,
				Desc: "5err-1",
			},
			{
				Id:   5,
				Desc: "5err-1",
			},
		}
	)

	assert.Nil(t, exist, "init")
	assert.Empty(t, exist.Error(), "init")

	assert.Nil(t, existResp.Errors, "init")

	exist.AddError(1, "err-", true)
	existResp.AddError(1, "err-", true)

	exist.AddErrorf(2, "err-%d", 15)
	existResp.AddErrorf(2, "err-%d", 15)

	errList := []error{
		errors.New("3err-1"),
		errors.New("3err-2"),
	}
	errRecList := RespErrorRecordList{
		{
			Id:   4,
			Desc: "5err-1",
		},
		{
			Id:   5,
			Desc: "5err-1",
		},
	}

	exist.AppendError(3, errList)
	existResp.AddError(3, errList[0].Error())
	existResp.AddError(3, errList[1].Error())

	exist.Append(errRecList)
	existResp.AppendError(errRecList)

	assert.Equal(t, expected, exist)
	assert.Equal(t, expected, existResp.Errors)

	assert.Error(t, exist.Result(), "checking an error")

	assert.True(t, RespErrorRecordList{}.Empty(), "check empty error list")
	assert.NoError(t, RespErrorRecordList{}.Result(), "check empty error")

	expectedStr := "[1] err-true; [2] err-15; [3] 3err-1; [3] 3err-2; [4] 5err-1; [5] 5err-1"
	assert.Equal(t, expectedStr, exist.Error())
}

func TestListOfErr_Error(t *testing.T) {
	var (
		exist = RespListOfErr{
			Id: 11,
		}
	)

	assert.Equal(t, false, exist.HasError(), "init")

	exist.Check(false, "err1")
	exist.Check(true, "err2")
	exist.Checkf(false, "err%d", 3)
	exist.Checkf(true, "err%d", 4)

	assert.Equal(t, true, exist.HasError())

	expectedStr := "[11] err2; [11] err4"
	assert.Equal(t, expectedStr, exist.Error())
}

func TestRespListOfErr(t *testing.T) {
	assert.NoError(t, (&RespListOfErr{}).Result(), "no one error")
	assert.False(t, (&RespListOfErr{}).HasError(), "no one error")

	respErr := &RespListOfErr{
		Id: errCode.ErrParamBinding,
	}
	respErr.Check(1 != 1, "err2")
	assert.NoError(t, respErr.Result(), "no one error")
	assert.False(t, respErr.HasError(), "no one error")

	respErr.Check(1 != 2, "err1")
	respErr.Checkf("122" != "233", "err2: %s", os.ErrInvalid)

	assert.Error(t, respErr.Result(), "a list of error")
	assert.True(t, respErr.HasError(), "a list of error")
	assert.Equal(t, fmt.Sprintf("[%d] err1; [%d] err2: %v", errCode.ErrParamBinding, errCode.ErrParamBinding, os.ErrInvalid), respErr.Error())
}

func TestNewGeneralErrorResponse(t *testing.T) {
	err := NewGeneralErrorResponse(errCode.ErrParamValidation, os.ErrInvalid)

	assert.Equal(t, &GeneralErrorResponse{
		RespError: RespError{
			Errors: RespErrorRecordList{
				{Id: errCode.ErrParamValidation, Desc: fmt.Sprint(os.ErrInvalid)},
			},
		},
	}, err)

	err = NewGeneralErrorResponse(errCode.ErrParamValidation, "err1")

	assert.Equal(t, &GeneralErrorResponse{
		RespError: RespError{
			Errors: RespErrorRecordList{
				{Id: errCode.ErrParamValidation, Desc: "err1"},
			},
		},
	}, err)

	err = NewGeneralErrorResponse(errCode.ErrParamValidation, RespErrorRecordList{
		{Id: errCode.ErrParamBinding, Desc: "err2"},
		{Id: errCode.ErrParamValidation, Desc: "err3"},
	})

	assert.Equal(t, &GeneralErrorResponse{
		RespError: RespError{
			Errors: RespErrorRecordList{
				{Id: errCode.ErrParamBinding, Desc: "err2"},
				{Id: errCode.ErrParamValidation, Desc: "err3"},
			},
		},
	}, err)
}
