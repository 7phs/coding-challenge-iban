package handler

import (
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/helper"
	"github.com/7phs/coding-challenge-iban/model"
	"github.com/7phs/coding-challenge-iban/model/errCode"
	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/7phs/coding-challenge-iban/model/validator"
	"github.com/7phs/coding-challenge-iban/restapi/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ValidateHandlerResponse struct {
	helper.RespError
	Data validator.Status `json:"data"`
}

type ValidateHandler struct {
	request struct {
		iban string
	}
	response struct {
		Status validator.Status `json:"iban"`
	}
}

func (o *ValidateHandler) New() common.HandlerImpl {
	return &ValidateHandler{}
}

func (o *ValidateHandler) LogPrefix() string {
	return "validate"
}

func (o *ValidateHandler) Bind(c *gin.Context) error {
	o.request.iban = strings.TrimSpace(c.Param("IBAN"))

	return nil
}

func (o *ValidateHandler) Validate(conf *config.Config) error {
	lst := helper.RespListOfErr{
		Id: errCode.ErrParamValidation,
	}

	if !lst.Check(len(o.request.iban) == 0, "iban: empty") {
		textLimit := conf.Limit().TextLength()

		lst.Check(textLimit > 0 && len(o.request.iban) > textLimit, "iban: longer than allowed")
	}

	return lst.Result()
}

func (o *ValidateHandler) Process(c *gin.Context) {
	if err := model.Validator().Validate(records.NewIban(o.request.iban)); err != nil {
		o.response.Status = validator.Invalid
	} else {
		o.response.Status = validator.Valid
	}

	c.JSON(http.StatusOK, o.response)
}