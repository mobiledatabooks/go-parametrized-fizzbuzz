package fizzbuzz

import (
	"testing"
)

// Copyright © 2022 Constantine Vassilev.

// func main() {
// 	var t *testing.T
// 	TestFuzzBuzzFunc(t)
// 	// FizzBuzzAllGen()
// }

var resultStr []string
var resultInt []int

// Variables used for test
//
// .variables
// [source,go]
// ----
// include::${gad:current:fq}[tag=fizz-buzz-test-variables,indent=0]
// ----
//
// tag::fizz-buzz-test-variables[]

// end::fizz-buzz-test-variables[]

// go test -run TestFuzzBuzzFunc01 -v
// go test -run=TestFuzzBuzzFunc01 -coverprofile=c.out
// go tool cover -html=c.out

// TestFuzzBuzzFunc01 performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-test-01[]
func TestFuzzBuzzFunc01(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			expected := FizzBuzz01(fizzAti, buzzAti) // <3>

			observed := FizzBuzz(fizzAti, buzzAti) // <4>

			// <5>
			// fmt.Println("===")
			// fmt.Println("observed")
			// fmt.Println(observed)
			// fmt.Println("expected")
			// fmt.Println(expected)
			// fmt.Println("===")

			if !equal(observed, expected) { // <6>
				t.Fatalf("\nFizzBuzz(total, fizzAt, buzzAt)) = \n%v, \nwant: \n%v\nfizzAti=%d buzzAti=%d\n", observed, expected, fizzAti, buzzAti)
			}
		}
	}
}

// end::fizz-buzz-test-01[]

// TestFuzzBuzzFunc01A performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-01-a[]
func TestFuzzBuzzFunc01A(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			expected := FizzBuzz01A(fizzAti, buzzAti) // <3>

			observed := FizzBuzzA(fizzAti, buzzAti) // <4>

			// <5>
			// fmt.Println("===")
			// fmt.Println("observed")
			// fmt.Println(observed)
			// fmt.Println("expected")
			// fmt.Println(expected)
			// fmt.Println("===")

			if !equal(observed, expected) { // <6>
				t.Fatalf("\nFizzBuzz(total, fizzAt, buzzAt)) = \n%v, \nwant: \n%v\nfizzAti=%d buzzAti=%d\n", observed, expected, fizzAti, buzzAti)
			}
		}
	}
}

// end::fizz-buzz-test-01-a[]

// TestFuzzBuzzFunc02 performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-02[]
func TestFuzzBuzzFunc02(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			expected := FizzBuzz02(fizzAti, buzzAti) // <3>

			observed := FizzBuzz(fizzAti, buzzAti) // <4>

			// <5>
			// fmt.Println("===")
			// fmt.Println("observed")
			// fmt.Println(observed)
			// fmt.Println("expected")
			// fmt.Println(expected)
			// fmt.Println("===")

			if !equal(observed, expected) { // <6>
				t.Fatalf("\nFizzBuzz(total, fizzAt, buzzAt)) = \n%v, \nwant: \n%v\nfizzAti=%d buzzAti=%d\n", observed, expected, fizzAti, buzzAti)
			}
		}
	}
}

// end::fizz-buzz-test-02[]

// TestFuzzBuzzFunc03 performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-03[]
func TestFuzzBuzzFunc03(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			expected := FizzBuzz03(fizzAti, buzzAti) // <3>

			observed := FizzBuzz(fizzAti, buzzAti) // <4>

			// <5>
			// fmt.Println("===")
			// fmt.Println("observed")
			// fmt.Println(observed)
			// fmt.Println("expected")
			// fmt.Println(expected)
			// fmt.Println("===")

			if !equal(observed, expected) { // <6>
				t.Fatalf("\nFizzBuzz(total, fizzAt, buzzAt)) = \n%v, \nwant: \n%v\nfizzAti=%d buzzAti=%d\n", observed, expected, fizzAti, buzzAti)
			}
		}
	}
}

// end::fizz-buzz-test-03[]

// TestFuzzBuzzFunc04 performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-04[]
func TestFuzzBuzzFunc04(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			expected := FizzBuzz04(fizzAti, buzzAti) // <3>

			observed := FizzBuzz(fizzAti, buzzAti) // <4>

			// <5>
			// fmt.Println("===")
			// fmt.Println("observed")
			// fmt.Println(observed)
			// fmt.Println("expected")
			// fmt.Println(expected)
			// fmt.Println("===")

			if !equal(observed, expected) { // <6>
				t.Fatalf("\nFizzBuzz(total, fizzAt, buzzAt)) = \n%v, \nwant: \n%v\nfizzAti=%d buzzAti=%d\n", observed, expected, fizzAti, buzzAti)
			}
		}
	}
}

// end::fizz-buzz-test-04[]

// BenchmarkFizzBuzz performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-benchmark[]
func benchmarkFizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				r = FizzBuzz(fizzAti, buzzAti) // <1>
			}
		}
	}
	resultStr = r
}

func BenchmarkFizzBuzz(b *testing.B) { benchmarkFizzBuzz(17, 17, b) }

// end::fizz-buzz-benchmark[]

// BenchmarkFizzBuzzA performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-a-benchmark[]
func benchmarkAFizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				r = FizzBuzzA(fizzAti, buzzAti) // <1>
			}
		}
	}
	resultStr = r
}
func BenchmarkAFizzBuzz(b *testing.B) { benchmarkAFizzBuzz(17, 17, b) }

// end::fizz-buzz-a-benchmark[]

// BenchmarkFizzBuzz01 performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-benchmark-01[]

func benchmark01FizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //				// always record the result of Fib to prevent
				r = FizzBuzz01(fizzAti, buzzAti) // <1>
			}
		}
	}
	resultStr = r
}

func Benchmark01FizzBuzz(b *testing.B) { benchmark01FizzBuzz(17, 17, b) }

// end::fizz-buzz-benchmark-01[]

// BenchmarkFizzBuzz02 performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-benchmark-02[]

func benchmark02FizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				r = FizzBuzz02(fizzAti, buzzAti) // <1>
			}
		}
	}
	resultStr = r
}

func Benchmark02FizzBuzz(b *testing.B) { benchmark02FizzBuzz(17, 17, b) }

// end::fizz-buzz-benchmark-02[]

// BenchmarkFizzBuzz03 performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-benchmark-03[]

func benchmark03FizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				r = FizzBuzz03(fizzAti, buzzAti) // <1>
			}
		}
	}
	resultStr = r
}

func Benchmark03FizzBuzz(b *testing.B) { benchmark02FizzBuzz(17, 17, b) }

// end::fizz-buzz-benchmark-03[]

// BenchmarkFizzBuzz04 performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-benchmark-04[]

func benchmark04FizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				r = FizzBuzz04(fizzAti, buzzAti) // <1>
			}
		}
	}
	resultStr = r
}

func Benchmark04FizzBuzz(b *testing.B) { benchmark04FizzBuzz(17, 17, b) }

// end::fizz-buzz-benchmark-04[]

// Check if two string arrays are equal.
// tag::equal[]
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// end::equal[]

// Check if two int arrays are equal.
// tag::equalint[]
func equalint(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// end::equalint[]
