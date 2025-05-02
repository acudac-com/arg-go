package arg

import "fmt"

type errs struct {
	msg string
}

// Will return first error encountered.
func Errors() *errs {
	return nil
}

func (c *errs) Add(invalid bool, msg string, args ...any) *errs {
	if c == nil {
		if invalid {
			c = &errs{fmt.Sprintf(msg, args...)}
		}
	}
	return c
}

func (c *errs) AddF(f func() error) *errs {
	if c == nil {
		if err := f(); err != nil {
			c = &errs{err.Error()}
		}
	}
	return c
}

func (c *errs) Error() string {
	if c == nil {
		return ""
	}
	return c.msg
}
