package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"

	"mobiledatabooks.com/go-parametrized-fizzbuzz/fizzbuzz"
)

// Copyright © 2022 Constantine Vassilev.

// go test -v -cover ./...

// The unit-testing contains a FizzBuzz example implemented in go. It takes three optional command line arguments.
//
var cpuprofile = flag.String("cpuprofile", "fizzbuzzCPUprofile.prof", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "fizzbuzzMemprofile.prof", "write memory profile to `file`")

// tag::fizz-buzz-main[]

// Main performs a FizzBuzz operation over a range of integers
func main() { // <1>
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	fmt.Println("FizzBuzz test")
	fmt.Println()

	// Set default argument values
	totalArgument := "20"
	fizzAtArgument := "3"
	buzzAtArgument := "5"

	// Convert argument values to numbers
	total, err := strconv.ParseInt(totalArgument, 10, 32)
	if err != nil {
		panic("The number of items to FizzBuzz should be an integer")
	}

	fizzAt, err := strconv.ParseInt(fizzAtArgument, 10, 32)
	if err != nil {
		panic("The number to Fizz at should be an integer")
	}

	buzzAt, err := strconv.ParseInt(buzzAtArgument, 10, 32)
	if err != nil {
		panic("The number to Buzz at should be an integer")
	}
	fmt.Println()
	fmt.Println(fizzAt)
	fmt.Println(buzzAt)
	// FizzBuzz the input and print the results
	fmt.Println(fmt.Sprintf("FizzBuzzing %d number(s), fizzing at %d and buzzing at %d:", total, fizzAt, buzzAt))
	for i, result := range fizzbuzz.FizzBuzz(fizzAt, buzzAt) {
		fmt.Printf("%2d %s\n", i+1, result)
	}

	fmt.Println()

	fmt.Println("Done")
	fmt.Println()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

// end::fizz-buzz-main[]
