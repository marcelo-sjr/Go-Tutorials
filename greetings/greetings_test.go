package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Alex"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Alex")
    if !want.MatchString(msg) || err != nil {
        t.Errorf(`Hello("Alex") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}


func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Errorf(`Hello("") expected "",error got %q %v`, message, err)
	}
}

func TestHellos(t *testing.T) {
	names := []string{"Jonas","Taylor","Bronn",}
	messages,err := Hellos(names)
	for _,name := range names{
		message,ok := messages[name]
		if !ok {
			t.Errorf("Key %s, was expected, but wasn't found.", name)
		}

		want:=regexp.MustCompile(`\b` +name+ `\b`)
		if !want.MatchString(messages[name]){
			t.Errorf(`messages[%s] = %s, want match for %#q`,name, message, want)
		}
	}
	if err != nil {
		t.Errorf("error = %v, want nil", err)
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"Alex","","Jonas"}
	messages, err := Hellos(names)
	if err == nil || len(messages) > 0 {
		t.Errorf(`Hellos(names) = %v, nil, want nil, error`,messages)
	}
}
