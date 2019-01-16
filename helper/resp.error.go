package helper

import (
	"bytes"
	"fmt"
)

type RespErrors interface {
	AddError(int, ...interface{})
	AddErrorf(int, string, ...interface{})
}

type RespErrorRecord struct {
	Id   int    `json:"id"`
	Desc string `json:"desc"`
}

type RespErrorRecordList []*RespErrorRecord

func (o *RespErrorRecordList) addError(id int, desc string) {
	*o = append(*o, &RespErrorRecord{Id: id, Desc: desc})
}

func (o *RespErrorRecordList) AddError(id int, v ...interface{}) {
	o.addError(id, fmt.Sprint(v...))
}

func (o *RespErrorRecordList) AddErrorf(id int, format string, v ...interface{}) {
	o.addError(id, fmt.Sprintf(format, v...))
}

func (o *RespErrorRecordList) Append(errList RespErrorRecordList) {
	*o = append(*o, errList...)
}

func (o *RespErrorRecordList) AppendError(id int, errList []error) {
	for _, err := range errList {
		o.AddError(id, err.Error())
	}
}

func (o RespErrorRecordList) Error() string {
	if len(o) == 0 {
		return ""
	}

	buf := bytes.NewBufferString("")

	for i, err := range o {
		if i > 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(fmt.Sprintf("[%d] %s", err.Id, err.Desc))
	}

	return buf.String()
}

func (o RespErrorRecordList) Empty() bool {
	return len(o) == 0
}

func (o RespErrorRecordList) Result() error {
	if o.Empty() {
		return nil
	}

	return o
}

type RespError struct {
	Errors RespErrorRecordList `json:"error"`
}

func (o *RespError) AddError(id int, v ...interface{}) {
	o.Errors.AddError(id, v...)
}

func (o *RespError) AddErrorf(id int, format string, v ...interface{}) {
	o.Errors.AddErrorf(id, format, v...)
}

func (o *RespError) AppendError(errList RespErrorRecordList) {
	o.Errors.Append(errList)
}

type RespListOfErr struct {
	Id int

	err RespErrorRecordList
}

func (o *RespListOfErr) Check(errCondition bool, errMessage string) bool {
	if errCondition {
		o.err.AddError(o.Id, errMessage)
	}

	return errCondition
}

func (o *RespListOfErr) Checkf(errCondition bool, format string, v ...interface{}) bool {
	if errCondition {
		o.err.AddErrorf(o.Id, format, v...)
	}

	return errCondition
}

func (o *RespListOfErr) HasError() bool {
	return o != nil && len(o.err) > 0
}

func (o *RespListOfErr) Error() string {
	return o.err.Error()
}

func (o *RespListOfErr) Result() error {
	if !o.HasError() {
		return nil
	}

	return o
}

type GeneralErrorResponse struct {
	RespError
}

func NewGeneralErrorResponse(id int, err interface{}) *GeneralErrorResponse {
	resp := &GeneralErrorResponse{}

	switch v := err.(type) {
	case RespErrorRecordList:
		resp.AppendError(v)
	case error:
		resp.AddError(id, v)
	default:
		resp.AddError(id, err)
	}

	return resp
}
