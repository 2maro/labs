package main

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Parrallel()
	c := counter.NewCounter()
	c.Input = bytes.NewBufferString("1\n2\n3\n4")
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}
