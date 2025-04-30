package arg_test

import (
	"testing"

	"github.com/acudac-com/arg-go"
)

func TestStringArg_Is(t *testing.T) {
	s := arg.S("test").Is("asdf", "qwer", "test")
	t.Logf("%v", s.Errors())
}
