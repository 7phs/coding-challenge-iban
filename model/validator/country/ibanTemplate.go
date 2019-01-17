package country

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseIbanTemplate(template string) (string, error) {
	if len(template) == 0 {
		return "", fmt.Errorf("empty template")
	}

	rg := bytes.NewBufferString("^")

	for _, part := range strings.Split(template, ",") {
		part = strings.TrimSpace(strings.ToLower(part))

		k, err := func(k string) (string, error) {
			switch k {
			case "a":
				return "[A-Z]", nil
			case "n":
				return "[0-9]", nil
			case "c":
				return "[a-zA-Z0-9]", nil
			default:
				return "", fmt.Errorf("unknown format character '%s'", k)
			}
		}(part[len(part)-1:])
		if err != nil {
			return "", err
		}

		n, err := strconv.Atoi(part[:len(part)-1])
		if err != nil {
			return "", err
		}

		rg.WriteString(fmt.Sprintf("%s{%d}", k, n))
	}

	rg.WriteString("$")

	return rg.String(), nil
}

type IbanTemplate struct {
	*regexp.Regexp
}

func NewIbanTemplate(template string) (*IbanTemplate, error) {
	return (&IbanTemplate{}).Parse(template)
}

func MustIbanTemplate(template string) *IbanTemplate {
	o, _ := (&IbanTemplate{}).Parse(template)

	return o
}

func (o *IbanTemplate) Nil() bool {
	return o.Regexp == nil
}

func (o *IbanTemplate) Parse(template string) (*IbanTemplate, error) {
	rg, err := ParseIbanTemplate(template)
	if err != nil {
		return nil, err
	}

	regexp, err := regexp.Compile(rg)
	if err != nil {
		return nil, err
	} else if regexp == nil {
		return nil, fmt.Errorf("failed to parse a template to regexp")
	}

	o.Regexp = regexp

	return o, nil
}

func (o *IbanTemplate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var template string

	if err := unmarshal(&template); err != nil {
		return err
	}

	if _, err := o.Parse(template); err != nil {
		return err
	}

	return nil
}
