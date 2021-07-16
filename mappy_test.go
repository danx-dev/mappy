package mappy_test

import (
	"encoding/json"
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

	result := mappy.DoMap(&input, &output)
	s, _ := json.Marshal(result)
	fmt.Println(string(s))

	fmt.Println(result)
}
