package arg_test

import (
	"testing"
	"time"

	"github.com/acudac-com/arg-go"
)

func TestCustom_FallbackIf(t *testing.T) {
	c := arg.C(time.Time{}).FallbackIf(time.Now(), false)
	t.Logf("%v", c.Value)
}
