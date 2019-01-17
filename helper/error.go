package helper

import "strings"

type ErrList []error

func (o *ErrList) Add(err error) {
	if err == nil {
		return
	}

	*o = append(*o, err)
}

func (o ErrList) Error() string {
	result := make([]string, 0, len(o))

	for _, err := range o {
		result = append(result, err.Error())
	}

	return strings.Join(result, "; ")
}

func (o ErrList) Result() error {
	if len(o) == 0 {
		return nil
	}

	return o
}
