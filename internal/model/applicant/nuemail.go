package applicant

import (
	"fmt"
	"regexp"
	"strings"
)

type NUEmail string

var re = regexp.MustCompile(`^[a-zA-Z]+\.[a-zA-Z]+[0-9]?@northeastern\.edu$`)

func ParseNUEmail(str string) (NUEmail, error) {
	if isNUEmail := re.MatchString(str); !isNUEmail {
		return "", fmt.Errorf("invalid northeastern email. got: %s", str)
	}

	return NUEmail(strings.ToLower(str)), nil
}

func (n *NUEmail) String() string {
	return string(*n)
}
