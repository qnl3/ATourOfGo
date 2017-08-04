package main

import (
	"fmt"
	."github.com/smartystreets/goconvey/convey"
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

func TestNeedInt(t *testing.T) {
	Convey("Untyped constants take on the type required by its context",t , func() {
		text := fmt.Sprintf("in a integer context small constanat renturs a int with the value of %d", 21)
		Convey(text, func(){
			x := needInt(small)
			So(x, ShouldEqual, 21)
		})
	})
}

func TestNeedFloat64(t *testing.T) {
	Convey("Untyped constants take on the type required by its context", t, func() {
		text0 := fmt.Sprintf("in a float context small constant returns a Float64 with the value of %d", 0.2)
		Convey(text0, func(){
			x := needFloat64(small)
			So(x, ShouldEqual, 0.2)
		})

		text1 := fmt.Sprintf("in a float context big constant returns a Float64 with the value of %d", 1.2676506002282295e+29)
		Convey(text1, func(){
			x := needFloat64(big)
			So(x, ShouldEqual, 1.2676506002282295e+29)
		})
	})
}

func ExampleNumericConstants() {
	numericConstants()
	// Output:
	// 21
	// 0.2
	// 1.2676506002282295e+29
}

func ExampleForLoop() {
	forLoop()
	// Output: 45
}

func ExampleForLoop2() {
	forLoop2()
	// Output: 1024
}

func ExampleForAsWhile(){
	forAsWhile()
	// Output: 1024
}

func TestSqrt(t *testing.T) {
	Convey("Given a 64 bit floating point number", t, func() {
		
		Convey("with a value of -4, sqrt(float64) returns 2i", func() {
			var input float64 = -4
			output := ifSqrt(input)
			So(output, ShouldEqual, "2i")
			
		})

		Convey("with a value of 2, sqrt(float64) returns 1.4142135623730951", func() {
			var input float64 = 2
			output := ifSqrt(input)
			So(output, ShouldEqual, "1.4142135623730951")
			
		})
	}) 
}

