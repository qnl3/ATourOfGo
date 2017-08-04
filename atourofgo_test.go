package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"

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
	var intValue int
	const x = (1 << 100) >> 99
	Convey(fmt.Sprintf("Given x is a integer with a value of %d", x), t, func() {
		Convey("When needInt(int) is call with x", func() {
			intValue = needInt(x)
			Convey(fmt.Sprintf("the functions interger return value(%d) should equal 21", intValue), func() {
				So(intValue, ShouldEqual, 21)
				So(intValue, ShouldHaveSameTypeAs, 21)
			})
		})
	})
}

func TestNeedFloat64(t *testing.T) {
	var f64Value float64
	const x = (1 << 100) >> 99
	var expected = 0.2
	Convey(fmt.Sprintf("Given x is an float64 with a value of %d", x), t, func() {
		Convey("When needfloat64(float64) is call with x", func() {
			f64Value = needFloat64(x)
			Convey("the functions returns a float64 value equal to 0.2", func() {
				So(f64Value, ShouldEqual, 0.2)
				So(f64Value, ShouldHaveSameTypeAs, expected)
			})
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

func ExampleForAsWhile() {
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
