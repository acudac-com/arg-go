package arg_test

import (
	"errors"
	"testing"

	"github.com/acudac-com/arg-go"
)

func TestErrors(t *testing.T) {
	country := "asdf"
	province := "sadf"
	street := "asdf"
	err := arg.Errors().
		Add(country == "", "invalid country").
		Add(province == "", "invalid province").
		AddF(func() error {
			if street == "" {
				return errors.New("invalid street")
			}
			return nil
		})
	if err != nil {
		logErr("address", err)
	}
}
