package main

import (
	. "github.com/smartystreets/goconvoy/convoy"
	"testing"
)

func ExampleHello() {
	hello()
	// Output: Hello, 世界!
}

func ExampleNamed() {
	named(17)
	// Output: 7 10
}

func ExampleVariables() {
	variables()
	// Output: 0 false false false
}

func ExampleVariableInt() {
	variableInit()
	// Output: 1 2 true false no!
}

func ExampleShortVar() {
	shortVar()
	// Output: 1 2 3 true false no!
}

func ExampleBasicTypes() {
	basicTypes()
	// Output:
	// Type: bool Value: false
	// Type: uint64 Value: 18446744073709551615
	// Type: complex128 Value: (2+3i)
}

func ExampleZero() {
	zero()
	// Output: 0 0 false ""
}

func ExampleTypeConversion() {
	typeConversion()
	// Output: 3 4 5
}

func ExampleTypeInferenceInt() {
	typeInferenceInt()
	// Output: v is of type int
}

func ExampleTypeInferenceFloat() {
	typeInferenceFloat64()
	// Output: v is of type float64
}

func ExampleTypeInferenceComplex128() {
	typeInferenceComplex128()
	// Output: v is of type complex128
}

func ExampleConstants() {
	constants()
	// Output:
	// Hello 世界
	// Happy 3.14 Day
	// Go rules? true
}

func ExampleNeedInt(t *testing.T) {
	needInt(2)
}

func ExampleNumericConstants() {
	numericConstants()
	// Output:
	// 21
	// 0.2
	// 1.2676506002282295e+29
}
