package records

import (
	"bytes"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

const (
	UnknownIban = "----"
)

var (
	cleaner = regexp.MustCompile("[^a-zA-Z0-9]+")
	replacer = regexp.MustCompile("[a-zA-Z]")
)

type Iban struct {
	raw string

	text string

	number big.Int
}

func NewIban(rawStr string) *Iban {
	return (&Iban{
		raw: rawStr,
	}).Parse(rawStr)
}

func (o *Iban) Parse(raw string) *Iban {
	o.text = o.clean(raw)

	if len(o.text) < 4 {
		o.text = UnknownIban
		return o
	}

	o.number = o.toNumber(o.text)

	return o
}

func (o *Iban) clean(raw string) string {
	return cleaner.ReplaceAllString(raw, "")
}

func (o *Iban) toNumber(str string) (number big.Int) {
	str = strings.ToUpper(str[4:]+str[0:4])

	str = replacer.ReplaceAllStringFunc(str, func(ch string) string {
		return strconv.Itoa(10 + int(ch[0]-'A'))
	})

	number.UnmarshalText([]byte(str))

	return
}

func (o *Iban) Len() int {
	return len(o.text)
}

func (o *Iban) Raw() string {
	return o.raw
}

func (o *Iban) Text() string {
	return o.text
}

func (o *Iban) Number() *big.Int {
	return &o.number
}

func (o *Iban) CountryCode() string {
	return o.text[0:2]
}

func (o *Iban) Kk() string {
	return o.text[2:4]
}

func (o *Iban) Suffix() string {
	return o.text[4:]
}

func (o *Iban) String() string {
	buf := bytes.NewBufferString("")
	i := 0

	for j:=4; j<len(o.text); j+=4 {
		buf.WriteString(o.text[i:j])
		buf.WriteString(" ")
		i+=4
	}
	// write a rest of the string
	if i<len(o.text) {
		buf.WriteString(o.text[i:])
	}

	return strings.TrimRight(buf.String(), " ")
}
