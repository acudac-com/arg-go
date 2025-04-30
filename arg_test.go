package arg_test

import (
	"testing"
	"time"

	"github.com/acudac-com/arg-go"
)

func TestArg_FallbackIf(t *testing.T) {
	c := arg.New(time.Time{}).FallbackIf(time.Now(), false)
	t.Logf("%v", c.Value)
	if c.Value.IsZero() {
		c.AddError("should not be zero")
	}
	t.Logf("valid: %t", c.Valid())
}

func Test_Invalid(t *testing.T) {
	c := arg.New(time.Time{})
	s := arg.S("asdf").StartsWith("order:")
	t.Logf("invalid: %t", arg.Invalid(c, s))
}
