package greetings

import (
	"regexp"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Alex"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Alex")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Alex") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Jonas", "Taylor", "Bronn"}
	messages, err := Hellos(names)
	if err != nil {
		t.Fatalf("error = %v, want nil", err)
	}

	for _, name := range names {
		message, ok := messages[name]
		if !ok {
			t.Fatalf("expected key %q to exist", name)
		}

		if !strings.Contains(message, name) {
			t.Fatalf(`message %q does not contain %q`, message, name)
		}
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"Alex", "", "Jonas"}
	messages, err := Hellos(names)
	if err == nil || len(messages) > 0 {
		t.Errorf(`Hellos(names) = %v, nil, want nil, error`, messages)
	}
}
