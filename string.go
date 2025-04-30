package arg

import (
	"regexp"
	"strings"
)

const (
	emailRgx  = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	domainRgx = `^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,}$`
	urlRegex  = `^(http://|https://|www\.)[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
)

// A string argument
type StringArg struct {
	*Arg[string]
}

// Returns a string argument instance for the provided value.
func String(value string) *StringArg {
	custom := &Arg[string]{value, []string{}}
	return &StringArg{custom}
}

// Returns a string argument instance for the provided value.
func S(value string) *StringArg {
	return String(value)
}

// Sets the value to the specified value if it's empty.
func (a *StringArg) FallbackIfEmpty(fallback string) *StringArg {
	if a.Value == "" {
		a.Value = fallback
	}
	return a
}

// Adds an error if the string is empty.
func (a *StringArg) Populated() *StringArg {
	if a.Value == "" {
		a.AddError("must be populated")
	}
	return a
}

// Adds an error if the string is populated.
func (a *StringArg) Empty() *StringArg {
	if a.Value != "" {
		a.AddError("must be empty")
	}
	return a
}

// Adds an error if the string is not equal to one of the specified values.
func (a *StringArg) Is(values ...string) *StringArg {
	for _, val := range values {
		if val == a.Value {
			return a
		}
	}
	if len(values) == 1 {
		a.AddError("must be %s", values[0])
	} else if len(values) > 1 {
		a.AddError("must be one of %s", strings.Join(values, ", "))
	}
	return a
}

// Adds an error if the string equals one of the specified values.
func (a *StringArg) IsNot(values ...string) *StringArg {
	for _, val := range values {
		if val == a.Value {
			a.AddError("must not be one of %s", strings.Join(values, ", "))
			return a
		}
	}
	return a
}

// Adds an error if the string does start with one of the specified prefixes.
func (a *StringArg) StartsWith(prefixes ...string) *StringArg {
	for _, prefix := range prefixes {
		if strings.HasPrefix(a.Value, prefix) {
			return a
		}
	}
	if len(prefixes) == 1 {
		a.AddError("must start with %s", prefixes[0])
	} else if len(prefixes) > 1 {
		a.AddError("must start with one of %s", strings.Join(prefixes, ", "))
	}
	return a
}

// Adds an error if the string does not end with one of the specified prefixes.
func (a *StringArg) EndsWith(suffixes ...string) *StringArg {
	for _, suffix := range suffixes {
		if strings.HasSuffix(a.Value, suffix) {
			return a
		}
	}
	if len(suffixes) == 1 {
		a.AddError("must end with %s", suffixes[0])
	} else if len(suffixes) > 1 {
		a.AddError("must end with one of %s", strings.Join(suffixes, ", "))
	}
	return a
}

// Adds an error if the string does not contain the specified substring.
func (a *StringArg) Contains(substring string) *StringArg {
	if strings.Contains(a.Value, substring) {
		return a
	}
	a.AddError("must contain %s", substring)
	return a
}

// Adds an error if the string does not match the specified regex.
func (a *StringArg) Matches(regex string) *StringArg {
	if !regexp.MustCompile(regex).MatchString(a.Value) {
		a.AddError("must match %s", regex)
	}
	return a
}

// Adds an error if the string is not a valid email
func (a *StringArg) IsEmail() *StringArg {
	if !regexp.MustCompile(emailRgx).MatchString(a.Value) {
		a.AddError("must be a valid email address")
	}
	return a
}

// Adds an error if the string is not a valid email and not empty.
func (a *StringArg) IsEmailOrEmpty() *StringArg {
	if a.Value != "" && !regexp.MustCompile(emailRgx).MatchString(a.Value) {
		a.AddError("must be a valid email address if specified")
	}
	return a
}

// Adds an error if the string is not a valid domain
func (a *StringArg) IsDomain() *StringArg {
	if !regexp.MustCompile(domainRgx).MatchString(a.Value) {
		a.AddError("must be a valid domain")
	}
	return a
}

// Adds an error if the string is not a valid domain and not empty.
func (a *StringArg) IsDomainOrEmpty() *StringArg {
	if a.Value != "" && !regexp.MustCompile(domainRgx).MatchString(a.Value) {
		a.AddError("must be a valid domain if specified")
	}
	return a
}

// Adds an error if the string is not a valid URL
func (a *StringArg) IsUrl() *StringArg {
	if !regexp.MustCompile(urlRegex).MatchString(a.Value) {
		a.AddError("must be a valid URL")
	}
	return a
}

// Adds an error if the string is not a valid URL and not empty.
func (a *StringArg) IsUrlOrEmpty() *StringArg {
	if a.Value != "" && !regexp.MustCompile(urlRegex).MatchString(a.Value) {
		a.AddError("must be a valid URL if specified")
	}
	return a
}

// Adds an error if the string is empty or its length is more than 60 characters
func (a *StringArg) IsTitle() *StringArg {
	if a.Value == "" || len(a.Value) > 60 {
		a.AddError("must be populated and no more than 60 characters")
	}
	return a
}

// Adds an error if the string's length is more than 120 characters
// Does not add error if empty.
func (a *StringArg) IsSubtitle() *StringArg {
	if a.Value != "" && len(a.Value) > 120 {
		a.AddError("must be no more than 120 characters if specified")
	}
	return a
}

// Adds an error if the string's length is more than 1000 characters
// Does not add error if empty.
func (a *StringArg) IsDescription() *StringArg {
	if a.Value != "" && len(a.Value) > 1000 {
		a.AddError("must be no more than 1000 characters if specified")
	}
	return a
}

// Adds an error if the string's length is not in the specified range
func (a *StringArg) LengthInRange(min, max int) *StringArg {
	l := len(a.Value)
	if l < min || l > max {
		a.AddError("must be between %d and %d characters", min, max)
	}
	return a
}
