package arg_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/acudac-com/arg-go"
)

func log(thing string, object any) {
	println(fmt.Sprintf("\033[1;34m%s\033[0m: %v", thing, object))
}

func logErr(thing string, object any) {
	println(fmt.Sprintf("\033[1;31m%s error\033[0m: %v", thing, object))
}

func TestArg_FallbackIf(t *testing.T) {
	x := &time.Time{}
	c := arg.New(x).FallbackIf(time.Now(), true)
	log("x", x)
	log("valid", c.Valid())
}
