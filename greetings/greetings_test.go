package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Manuel"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Manuel") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)

	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{
		"Manuel",
		"Javier",
		"Jaime",
	}

	messages, err := Hellos(names)

	if err != nil {
		t.Errorf("Not error expected, should return a map of differents greetings.")
	}

	rightKeys := true
	for _, name := range names {
		value, ok := messages[name]
		if !ok || value == "" {
			rightKeys = false
		}
	}

	if len(messages) == 0 {
		t.Errorf("Messages len should be more than 0")
	}

	if !rightKeys {
		t.Errorf("Messages should content the right keys")
	}

	if len(messages) != len(names) {
		t.Errorf("Messages len should be the same as names len")
	}
}
