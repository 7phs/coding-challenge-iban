package validator

import (
	"math/big"
	"os"

	"github.com/7phs/coding-challenge-iban/model/records"
)

var (
	modBase         = big.NewInt(97)
	modBaseCheckSum = big.NewInt(1)
)

type CheckSum struct{}

func NewCheckSum() *CheckSum {
	return &CheckSum{}
}

func (o *CheckSum) Validate(rec *records.Iban) error {
	if modBaseCheckSum.Cmp((&big.Int{}).Mod(rec.Number(), modBase)) != 0 {
		return os.ErrInvalid
	}

	return nil
}
