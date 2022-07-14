package greetings

import (
  "testing"
  "regexp"
)

func TestHelloName(t *testing.T){
  name := "Meth"
  want := regexp.MustCompile(`\b`+name+`\b`)
  msg, err := Hello("Meth")
  if !want.MatchString(msg) || err != nil {
    t.FatalF(`Hello("Meth") = %q, %v, want match for %#q, nil`, msg, err, want)
  }
}

func TestHelloEmpty(t *testing.T) {
  msg, err := Hello("")
  if msg != "" || err == nil {
    t.FatalF(`Hello("") = %q, %v, want "", error`, message, err)
  }
}
