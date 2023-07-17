package main

import (
	"fmt"
	"testing"
)

func Test_toSet(t *testing.T) {
	input1 := "dWlhclDHdFvDCCDfFq"
	set1 := toSet(input1)
	fmt.Printf("%v\n", set1)
}
