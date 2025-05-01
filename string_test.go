package arg_test

import (
	"testing"

	"github.com/acudac-com/arg-go"
)

func TestStringArg_Is(t *testing.T) {
	test := "test"
	s := arg.S(&test).Is("asdf", "qwer", "tst")
	log("valid", s.Valid())
}
