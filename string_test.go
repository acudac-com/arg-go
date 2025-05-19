package arg_test

import (
	"testing"

	"github.com/acudac-com/arg-go"
)

func TestStringArg_Is(t *testing.T) {
	test := "jan@alisx123.com"
	s := arg.S(&test).IsEmailWithExistingMx()
	log("valid", s.Valid())
}
