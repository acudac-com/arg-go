package arg_test

import (
	"errors"
	"testing"

	"github.com/acudac-com/arg-go"
)

type Address struct {
	Country  string
	Province string
	street   string
}

func (a *Address) Street() string {
	if a == nil {
		return ""
	}
	return a.street
}

func (a *Address) ValidCountry() error {
	if len(a.Country) < 2 {
		return errors.New("at least 2 country characters required")
	}
	return nil
}

func (a *Address) ValidProvince() error {
	if a.Province == "" {
		return errors.New("province required")
	}
	return nil
}

func TestErrors(t *testing.T) {
	addr := &Address{"ab", "asdf", "qwer"}
	if err := arg.Errors(addr == nil, "address may not be nil").
		Add(addr.ValidCountry()).
		AddF(addr.ValidProvince).
		AddB(addr.Street() == "", "invalid street"); err != nil {
		logErr("address", err)
	}
}
