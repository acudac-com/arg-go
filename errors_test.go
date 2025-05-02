package arg_test

import (
	"testing"

	"github.com/acudac-com/arg-go"
)

func TestErrors(t *testing.T) {
	country := ""
	province := "sadf"
	street := "asdf"
	if err := arg.Errors(country == "", "invalid country").
		Add(province == "", "invalid province").
		Add(street == "", "invalid street"); err != nil {
		logErr("address", err)
	}
}
