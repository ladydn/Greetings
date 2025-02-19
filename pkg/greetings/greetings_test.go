package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Lady"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Lady")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Lady") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
