package fizzbuzz

import (
	"strconv"
)

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
// 1
// 2
// Fizz
// 4
// Buzz
// Fizz
// 7
// 8
// Fizz
// Buzz
// 11
// Fizz
// 13
// 14
// FizzBuzz
// 16

// FizzBuzzSlow generates the expected array in an alternative way (very inefficient).
//
// tag::fizz-buzz-01[]
func FizzBuzzSlow(fizzAti, buzzAti int64) []string {
	var total int64 = 4 * fizzAti * buzzAti // <1>
	if buzzAti < fizzAti {                  // <2>
		t := buzzAti
		buzzAti = fizzAti
		fizzAti = t
	}
	result := make([]string, total) // <3>
	if buzzAti == fizzAti {         // <4>
		for i := int64(0); i < total; i++ {
			m := (i + 1) % buzzAti
			if m == 0 {
				result[i] = "Fizz" + "Buzz"
			} else {
				result[i] = strconv.FormatInt(i+1, 10) //Digit
			}
		}

	} else if buzzAti > fizzAti { // <5>
		for i := int64(0); i < total; i++ {
			m := (i + 1) % buzzAti
			n := (i + 1) % fizzAti

			if m == 0 { // <6>
				result[i] = "Fizz" + "Buzz"
			} else { // <7>
				result[i] = "Fizz"
			}

			for mi := int64(1); mi < total; mi++ { // <8>
				for ni := int64(1); ni < total; ni++ {
					if m == mi && n == ni {
						result[i] = strconv.FormatInt(i+1, 10) //Digit
					}
				}
			}
			for ni := int64(1); ni < total; ni++ { // <9>
				if m == 0 && n == ni {
					result[i] = "Buzz"
				}
			}
		}
	}
	return result
}

// end::fizz-buzz-01[]

// FizzBuzzSlowMinMemory generates the expected array in an alternative way (very inefficient).
//
// tag::fizz-buzz-01-a[]
func FizzBuzzSlowMinMemory(fizzAti, buzzAti int64) []string {
	var total int64 = 4 * fizzAti * buzzAti // <1>
	if buzzAti < fizzAti {                  // <2>
		t := buzzAti
		buzzAti = fizzAti
		fizzAti = t
	}
	result := make([]string, total) // <3>

	if buzzAti == fizzAti { // <4>
		for i := int64(0); i < total; i++ {
			m := (i + 1) % buzzAti
			if m == 0 {
				result[i] = Fizz + Buzz
			} else {
				result[i] = Digit // strconv.FormatInt(i+1, 10) //Digit
			}
		}

	} else if buzzAti > fizzAti { // <5>
		for i := int64(0); i < total; i++ {
			m := (i + 1) % buzzAti
			n := (i + 1) % fizzAti

			if m == 0 { // <6>
				result[i] = Fizz + Buzz
			} else { // <7>
				result[i] = Fizz
			}

			for mi := int64(1); mi < total; mi++ { // <8>
				for ni := int64(1); ni < total; ni++ {
					if m == mi && n == ni {
						result[i] = Digit // strconv.FormatInt(i+1, 10) //Digit
					}
				}
			}
			for ni := int64(1); ni < total; ni++ { // <9>
				if m == 0 && n == ni {
					result[i] = Buzz
				}
			}
		}
	}
	return result
}

// end::fizz-buzz-01-a[]

// FizzBuzzIfElse performs a FizzBuzz operation over a range of integers
// tag::fizz-buzz-02[]
func FizzBuzzIfElse(fizzAt, buzzAt int64) []string {
	var total int64 = 4 * fizzAt * buzzAt // <1>
	if buzzAt < fizzAt {                  // <2>
		t := buzzAt
		buzzAt = fizzAt
		fizzAt = t
	}
	result := make([]string, total) // <3>
	for i := int64(1); i <= total; i++ {

		// if i%(fizzAt*buzzAt) == 0 { // <4>
		// 	fmt.Println()
		// }
		// <5>
		if i%fizzAt == 0 && i%buzzAt == 0 { // <6>
			result[i-1] = "FizzBuzz"
		} else if i%fizzAt == 0 { // <7>
			result[i-1] = "Fizz"
		} else if i%buzzAt == 0 { //<8>
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.FormatInt(i, 10) // <9>
		}
		// fmt.Printf("result[%2d]==%10s \t(i=%2d) i%%fizzAt(%2d]==%2d i%%buzzAt(%2d]==%2d)\n", // <10>
		// 	i-1, result[i-1], i, fizzAt, i%fizzAt, buzzAt, i%buzzAt)

	}
	return result
}

// end::fizz-buzz-02[]

// FizzBuzzContinue performs a FizzBuzz operation over a range of integers
// tag::fizz-buzz-03[]
func FizzBuzzContinue(fizzAt, buzzAt int64) []string {
	var total int64 = 4 * fizzAt * buzzAt // <1>
	if buzzAt < fizzAt {                  // <2>
		t := buzzAt
		buzzAt = fizzAt
		fizzAt = t
	}
	result := make([]string, total) // <3>
	for i := int64(1); i <= total; i++ {

		// if i%(fizzAt*buzzAt) == 0 { // <4>
		// 	fmt.Println()
		// }

		if !(i%fizzAt == 0) && !(i%buzzAt == 0) { // <5>
			result[i-1] = strconv.FormatInt(i, 10)
			continue
		}

		if i%fizzAt == 0 { // <6>
			result[i-1] = "Fizz"
		}

		if i%buzzAt == 0 { // <7>
			result[i-1] += "Buzz"
		}

		// fmt.Printf("result[%2d]==%10s \t(i=%2d) i%%fizzAt(%2d]==%2d i%%buzzAt(%2d]==%2d)\n", // <8>
		// 	i-1, result[i-1], i, fizzAt, i%fizzAt, buzzAt, i%buzzAt)

	}
	return result
}

// end::fizz-buzz-03[]

// FizzBuzzBigInt performs a FizzBuzz operation over a range of integers

// tag::fizz-buzz-04[]
func FizzBuzzBigInt(fizzAt, buzzAt int64) []string {
	var total int64 = 4 * fizzAt * buzzAt // <1>
	if buzzAt < fizzAt {                  // <2>
		t := buzzAt
		buzzAt = fizzAt
		fizzAt = t
	}
	result := make([]string, total) // <3>
	for i := int64(1); i <= total; i++ {

		// if i%(fizzAt*buzzAt) == 0 { // <4>
		// 	fmt.Println()
		// }

		if i%fizzAt == 0 { // <6>
			result[i-1] = "Fizz"
		}

		if i%buzzAt == 0 { // <7>
			result[i-1] += "Buzz"
		}
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) { // <5>
			result[i-1] = strconv.FormatInt(i, 10)
			// continue
		}
		// fmt.Printf("result[%2d]==%10s \t(i=%2d) i%%fizzAt(%2d]==%2d i%%buzzAt(%2d]==%2d)\n", // <8>
		// 	i-1, result[i-1], i, fizzAt, i%fizzAt, buzzAt, i%buzzAt)

	}
	return result
}

// end::fizz-buzz-04[]

// FizzBuzz performs a FizzBuzz operation over a range of integers
// tag::fizz-buzz[]
func FizzBuzz(fizzAt, buzzAt int64) []string {
	var total int64 = 4 * fizzAt * buzzAt // <1>
	if buzzAt < fizzAt {                  // <2>
		t := buzzAt
		buzzAt = fizzAt
		fizzAt = t
	}
	result := make([]string, total) // <3>
	for i := int64(1); i <= total; i++ {

		// if i%(fizzAt*buzzAt) == 0 { // <4>
		// 	fmt.Println()
		// }

		if i%fizzAt == 0 { // <5>
			result[i-1] = "Fizz"
		}

		if i%buzzAt == 0 { // <6>
			result[i-1] += "Buzz"
		}

		if result[i-1] == "" { // <7>
			result[i-1] = strconv.FormatInt(i, 10)
		}

		// fmt.Printf("result[%2d]==%10s \t(i=%2d) i%%fizzAt(%2d]==%2d i%%buzzAt(%2d]==%2d)\n", // <8>
		// 	i-1, result[i-1], i, fizzAt, i%fizzAt, buzzAt, i%buzzAt)

	}
	return result
}

// end::fizz-buzz[]

const Fizz = "X" // "X"
const Buzz = "Y" // "Y"
// const FizzBuzz = "Z" //"Z"
const Digit = "O"

// FizzBuzzMinMemory performs a FizzBuzz operation over a range of integers
// tag::fizz-buzz-a[]
func FizzBuzzMinMemory(fizzAt, buzzAt int64) []string {
	var total int64 = 4 * fizzAt * buzzAt // <1>
	if buzzAt < fizzAt {                  // <2>
		t := buzzAt
		buzzAt = fizzAt
		fizzAt = t
	}
	result := make([]string, total) // <3>
	for i := int64(1); i <= total; i++ {

		// if i%(fizzAt*buzzAt) == 0 { // <4>
		// 	fmt.Println()
		// }

		if i%fizzAt == 0 { // <5>
			result[i-1] = "X" //"Fizz" //"X" //Fizz // Fizz
		}

		if i%buzzAt == 0 { // <6>
			result[i-1] += "Y" //"Buzz" //"Y" //Buzz // 2=Buzz 3=FizzBuzz
		}

		if result[i-1] == "" { // <7>
			result[i-1] = "O" //Digit // Digit strconv.FormatInt(i, 10)
		}

		// fmt.Printf("result[%2d]==%10s \t(i=%2d) i%%fizzAt(%2d]==%2d i%%buzzAt(%2d]==%2d)\n", // <8>
		// 	i-1, result[i-1], i, fizzAt, i%fizzAt, buzzAt, i%buzzAt)

	}
	return result
}

// end::fizz-buzz-a[]
const empty = "" //"x"
// FizzBuzzConcurrent  Concurrent FizzBuzz with Go Routines from Russ Cox. Here you find a parametrized solution.
//
// .FizzBuzzConcurrent
// tag::fizz-buzz-FizzBuzzConcurrent[]
func FizzBuzzConcurrent(fizzAt, buzzAt int64) []string { // <1>
	if buzzAt < fizzAt {
		t := buzzAt
		buzzAt = fizzAt
		fizzAt = t
	}
	c := generate() // <2>
	// <3>
	c = filter(c, int(fizzAt), "Fizz")    // <6>
	c = filter(c, int(buzzAt), "Buzz")    // <6>
	var total int64 = 4 * fizzAt * buzzAt //
	result := make([]string, total)       // <5>

	for i := int64(1); i <= total; i++ { // <7>
		if s := <-c; s != empty {
			result[i-1] = s

		} else {
			result[i-1] = strconv.FormatInt(i, 10) //
		}
	}
	return result // <4>
}

func generate() <-chan string { // <8>
	c := make(chan string)
	go func() { // <9>
		for { //i0 := 1; ; i0++ {
			c <- empty //+ " " + strconv.Itoa(i0) + " "
		}
	}()
	return c
}

// <10>
func filter(c <-chan string, n int, label string) <-chan string { // <3>
	out := make(chan string)
	go func() { // <12>
		for { //

			var i int
			for i = 1; i < n; i++ { // <13>
				out <- <-c
			}
			out <- <-c + label // <14>
		}
	}()
	return out // <11>
}

// end::fizz-buzz-FizzBuzzConcurrent[]
