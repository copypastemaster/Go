package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Enoch")

	got := buffer.String()
	want := "Hello, Enoch"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	} else {
		fmt.Println("Hey")
	}
}
