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

// TestFuzzBuzzFuncSlow performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-test-01[]
func TestFuzzBuzzFuncSlow(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			observed := FizzBuzzSlow(fizzAti, buzzAti) // <3>

			expected := FizzBuzz(fizzAti, buzzAti) // <4>

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

// TestFizzBuzzFuncSlowMinMemory performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-01-a[]
func TestFizzBuzzFuncSlowMinMemory(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			observed := FizzBuzzSlowMinMemory(fizzAti, buzzAti) // <3>

			expected := FizzBuzzMinMemory(fizzAti, buzzAti) // <4>

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

// TestFuzzBuzzFuncIfElse performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-02[]
func TestFuzzBuzzFuncIfElse(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			observed := FizzBuzzIfElse(fizzAti, buzzAti) // <3>

			expected := FizzBuzz(fizzAti, buzzAti) // <4>

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

// TestFuzzBuzzFuncContinue performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-03[]
func TestFuzzBuzzFuncContinue(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			observed := FizzBuzzContinue(fizzAti, buzzAti) // <3>

			expected := FizzBuzz(fizzAti, buzzAti) // <4>

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

// TestFuzzBuzzFuncBigInt performs a FizzBuzz test operation over all possible combinations fizzAt between buzzAt
// tag::fizz-buzz-test-04[]
func TestFuzzBuzzFuncBigInt(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17

	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			observed := FizzBuzzBigInt(fizzAti, buzzAti) // <3>

			expected := FizzBuzz(fizzAti, buzzAti) // <4>

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

// TestFuzzBuzzFuncConcurrent performs a concurrent FizzBuzz test over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-test-05[]

func TestFuzzBuzzFuncConcurrent(t *testing.T) {
	// fmt.Printf("Testing all possible combinations between fizzAt=%2d:buzzAt=%2d\n", fizzAt, buzzAt)
	var fizzAt int64 = 17 //2 3 17
	var buzzAt int64 = 17 //3 4 17
	// f, _ := os.Create("observed.txt")
	// defer f.Close()
	for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { // <1>
		for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { // <2>
			// fmt.Printf("fizzAti=%2d buzzAti=%2d\n",
			// 	fizzAti, buzzAti)

			observed := FizzBuzzConcurrent(fizzAti, buzzAti) // <3>
			// _, _ = f.WriteString(fmt.Sprintf("%v", observed))

			expected := FizzBuzz(fizzAti, buzzAti) // <4>

			// <5>
			// fmt.Println("===")
			// fmt.Println("observed")
			// fmt.Println(observed)
			// fmt.Println("expected")
			// fmt.Println(expected)
			// fmt.Println("===")

			if !equal(observed, expected) { // <6>
				t.Fatalf("\nFizzBuzzConcurrent(total, fizzAt, buzzAt)) = \n%v, \nwant: \n%v\nfizzAti=%d buzzAti=%d\n", observed, expected, fizzAti, buzzAti)
			}
		}
	}
	// observed, _ := ioutil.ReadFile("observed.txt")
	// expected, _ := ioutil.ReadFile("expected.txt")
	// if string(observed) != string(expected) { // <6>
	// 	t.Fatalf("\nFuzzBuzz  = \n%v, \nwant: \n%v\n", observed, expected)
	// }
}

// end::fizz-buzz-test-05[]

// BenchmarkFizzBuzz performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-benchmark[]

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func benchmarkProcessFizzBuzz(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				// log.Println(buzzAti)

				r = FizzBuzz(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}

func BenchmarkFizzBuzz(b *testing.B) { benchmarkProcessFizzBuzz(17, 17, b) }

// end::fizz-buzz-benchmark[]

// BenchmarkFizzBuzzSlowMinMemory performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-a-benchmark[]
func benchmarkProcessFizzBuzzSlowMinMemory(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = FizzBuzzSlowMinMemory(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}
func BenchmarkFizzBuzzSlowMinMemory(b *testing.B) { benchmarkProcessFizzBuzzSlowMinMemory(17, 17, b) }

// end::fizz-buzz-a-benchmark[]

// benchmarkFizzBuzzSlow performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-benchmark-01[]

func benchmarkProcessFizzBuzzSlow(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = FizzBuzzSlow(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}

func BenchmarkFizzBuzzSlow(b *testing.B) {
	benchmarkProcessFizzBuzzSlow(17, 17, b)
}

// end::fizz-buzz-benchmark-01[]

// BenchmarkFizzBuzz02 performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-benchmark-02[]

func benchmarkProcessFizzBuzzIfElse(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = FizzBuzzIfElse(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}

func BenchmarkFizzBuzzIfElse(b *testing.B) { benchmarkProcessFizzBuzzIfElse(17, 17, b) }

// end::fizz-buzz-benchmark-02[]

// BenchmarkFizzBuzz03 performs a FizzBuzz benchmark operation over all possible combinations fizzAt between buzzAt
//
// tag::fizz-buzz-benchmark-03[]

func benchmarkProcessFizzBuzzContinue(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = FizzBuzzContinue(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}

func BenchmarkFizzBuzzContinue(b *testing.B) { benchmarkProcessFizzBuzzContinue(17, 17, b) }

// end::fizz-buzz-benchmark-03[]

func benchmarkProcessFizzBuzzBigInt(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = FizzBuzzBigInt(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}
func BenchmarkFizzBuzzBigInt(b *testing.B) { benchmarkProcessFizzBuzzBigInt(17, 17, b) }

func benchmarkProcessFizzBuzzConcurrent(fizzAt, buzzAt int64, b *testing.B) {
	var r []string
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for fizzAti := int64(1); fizzAti <= fizzAt; fizzAti++ { //
			for buzzAti := int64(1); buzzAti <= buzzAt; buzzAti++ { //
				// always record the result of Fib to prevent
				// the compiler eliminating the function call.
				r = FizzBuzzConcurrent(fizzAti, buzzAti) // <1>
			}
		}
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultStr = r
}
func BenchmarkFizzBuzzConcurrent(b *testing.B) { benchmarkProcessFizzBuzzConcurrent(17, 17, b) }

// Check if two string arrays are equal.
//
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
//
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
