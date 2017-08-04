package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	c, python, java bool
	i, j            int    = 1, 2
	toBe                   = false
	maxInt          uint64 = 1<<64 - 1
	z                      = cmplx.Sqrt(-5 + 12i)
)

const (
	pi    = 3.14
	big   = 1 << 100
	small = big >> 99
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func hello() {
	fmt.Println("Hello, 世界!")
}

func named(sum int) {
	fmt.Println(split(sum))
}

func variables() {
	var i int
	fmt.Println(i, c, python, java)
}

func variableInit() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func shortVar() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

func basicTypes() {
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversion() {
	x, y := 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}

func typeInferenceInt() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}

func typeInferenceFloat64() {
	v := 3.142
	fmt.Printf("v is of type %T\n", v)
}

func typeInferenceComplex128() {
	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
}

func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func needInt(x int) int { return x*10 + 1 }

func needFloat64(x float64) float64 {return x * 0.1}

func numericConstants() {
	fmt.Println(needInt(small))
	fmt.Println(needFloat64(small))
	fmt.Println(needFloat64(big))
}

func forLoop() {
	sum :=  0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forLoop2() {
	sum := 1
	for ;sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func forAsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func ifSqrt(x float64) string {
	if x < 0 {
		return ifSqrt(-x) +"i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	hello()
	named(17)
	variables()
	variableInit()
	shortVar()
	basicTypes()
	zero()
	typeConversion()
	typeInferenceInt()
	typeInferenceFloat64()
	typeInferenceComplex128()
	forLoop()
	forLoop2()
	forAsWhile()
	fmt.Println(ifSqrt(2), ifSqrt(-4))
}
