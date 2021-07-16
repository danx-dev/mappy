package mappy_test

import (
	"fmt"
	"testing"

	"github.com/danx-dev/mappy"
)

type S1 struct {
	Value1 string `mappy:"Value1"`
	Value2 string `mappy:"Sub.Value"`
}

type S2 struct {
	Value1 string
	Sub    Sub
}

type Sub struct {
	Value string
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

func TestDeepMapping(t *testing.T) {
	input := S1{Value1: "Test", Value2: "GoSubMapping"}
	output := S2{}

	mappy.DoMap(&output, &input)

	fmt.Println("Output", output)
	if output.Sub.Value != "GoSubMapping" {
		t.Error("Failed to Map Nested Struct")
	}
}
