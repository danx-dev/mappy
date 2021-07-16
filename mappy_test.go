package mappy_test

import (
	"fmt"
	"testing"

	"github.com/danx-dev/mappy"
)

type S1 struct {
	Value1 string `mappy:"Value1"`
}

type S2 struct {
	Value1 string
}

func TestSimpleMapping(t *testing.T) {
	input := S1{Value1: "Test"}
	output := S2{Value1: ""}

	mappy.DoMap(&output, &input)

	fmt.Println("Output", output)
	if output.Value1 != "Test" {
		t.Error("Value1 was not mapped!")
	}
}
