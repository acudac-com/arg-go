package arg

import (
	"net"
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
	*ComparableArg[string]
}

// Returns a string argument instance for the provided value.
func String(value *string) *StringArg {
	return &StringArg{C(value)}
}

// Returns a string argument instance for the provided value.
func S(value *string) *StringArg {
	return String(value)
}

// Adds an error if the string does start with one of the specified prefixes.
func (a *StringArg) StartsWith(prefixes ...string) *StringArg {
	for _, prefix := range prefixes {
		if strings.HasPrefix(*a.value, prefix) {
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
		if strings.HasSuffix(*a.value, suffix) {
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
	if strings.Contains(*a.value, substring) {
		return a
	}
	a.AddError("must contain %s", substring)
	return a
}

// Adds an error if the string does not match the specified regex.
func (a *StringArg) Matches(regex string) *StringArg {
	if !regexp.MustCompile(regex).MatchString(*a.value) {
		a.AddError("must match %s", regex)
	}
	return a
}

// Adds an error if the string is not a valid email. Rather use
// IsEmailWithExistingMx, unless you do not care about the email's existance.
func (a *StringArg) IsEmail() *StringArg {
	if !regexp.MustCompile(emailRgx).MatchString(*a.value) {
		a.AddError("must be a valid email address")
	}
	return a
}

// Adds an error if the email is not valid or tis provider's MX record cannot be
// found.
func (a *StringArg) IsEmailWithExistingMx() *StringArg {
	if !regexp.MustCompile(emailRgx).MatchString(*a.value) {
		a.AddError("must be a valid email address")
		return a
	}

	// find mx
	emailParts := strings.Split(*a.value, "@")
	domain := emailParts[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		a.AddError("email dns must have valid mx record")
	}
	return a
}

// Adds an error if the string is not a valid email and not empty.
func (a *StringArg) IsEmailOrEmpty() *StringArg {
	if *a.value != "" && !regexp.MustCompile(emailRgx).MatchString(*a.value) {
		a.AddError("must be a valid email address if specified")
	}
	return a
}

// Adds an error if the string is not a valid domain
func (a *StringArg) IsDomain() *StringArg {
	if !regexp.MustCompile(domainRgx).MatchString(*a.value) {
		a.AddError("must be a valid domain")
	}
	return a
}

// Adds an error if the string is not a valid domain and not empty.
func (a *StringArg) IsDomainOrEmpty() *StringArg {
	if *a.value != "" && !regexp.MustCompile(domainRgx).MatchString(*a.value) {
		a.AddError("must be a valid domain if specified")
	}
	return a
}

// Adds an error if the string is not a valid URL
func (a *StringArg) IsUrl() *StringArg {
	if !regexp.MustCompile(urlRegex).MatchString(*a.value) {
		a.AddError("must be a valid URL")
	}
	return a
}

// Adds an error if the string is not a valid URL and not empty.
func (a *StringArg) IsUrlOrEmpty() *StringArg {
	if *a.value != "" && !regexp.MustCompile(urlRegex).MatchString(*a.value) {
		a.AddError("must be a valid URL if specified")
	}
	return a
}

// Adds an error if the string is empty or its length is more than 60 characters
func (a *StringArg) IsTitle() *StringArg {
	if *a.value == "" || len(*a.value) > 60 {
		a.AddError("must be populated and no more than 60 characters")
	}
	return a
}

// Adds an error if the string's length is more than 120 characters
// Does not add error if empty.
func (a *StringArg) IsSubtitle() *StringArg {
	if *a.value != "" && len(*a.value) > 120 {
		a.AddError("must be no more than 120 characters if specified")
	}
	return a
}

// Adds an error if the string's length is more than 1000 characters
// Does not add error if empty.
func (a *StringArg) IsDescription() *StringArg {
	if *a.value != "" && len(*a.value) > 1000 {
		a.AddError("must be no more than 1000 characters if specified")
	}
	return a
}

// Adds an error if the string's length is not in the specified range
func (a *StringArg) LengthInRange(min, max int) *StringArg {
	l := len(*a.value)
	if l < min || l > max {
		a.AddError("must be between %d and %d characters", min, max)
	}
	return a
}
