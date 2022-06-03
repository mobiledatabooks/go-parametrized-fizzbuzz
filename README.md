# go-parametrized-fizzbuzz
Golang Solution for Parametrized Fizzbuzz. Unit Testing for All Possible Input Combinations: How to Build and Deploy Production-Grade Go (Golang) projects.

View on Amazon
[Golang Solution for Parametrized Fizzbuzz. Unit Testing for All Possible Input Combinations: How to Build and Deploy Production-Grade Go (Golang) projects.](https://www.amazon.com/dp/B09ZTD459Z)

# go\-parametrized\-fizzbuzz

```go
import "mobiledatabooks.com/go-parametrized-fizzbuzz"
```

## Index

- [Variables](<#variables>)
- [func main()](<#func-main>)


## Variables

The unit\-testing contains a FizzBuzz example implemented in go\. It takes three optional command line arguments\.

```go
var cpuprofile = flag.String("cpuprofile", "fizzbuzzCPUprofile.prof", "write cpu profile to `file`")
```

```go
var memprofile = flag.String("memprofile", "fizzbuzzMemprofile.prof", "write memory profile to `file`")
```

## func main

```go
func main()
```

Main performs a FizzBuzz operation over a range of integers


# fizzbuzz

```go
import "mobiledatabooks.com/go-parametrized-fizzbuzz/fizzbuzz"
```

## Index

- [Constants](<#constants>)
- [func FizzBuzz(fizzAt, buzzAt int64) []string](<#func-fizzbuzz>)
- [func FizzBuzz01(fizzAti, buzzAti int64) []string](<#func-fizzbuzz01>)
- [func FizzBuzz01A(fizzAti, buzzAti int64) []string](<#func-fizzbuzz01a>)
- [func FizzBuzz02(fizzAt, buzzAt int64) []string](<#func-fizzbuzz02>)
- [func FizzBuzz03(fizzAt, buzzAt int64) []string](<#func-fizzbuzz03>)
- [func FizzBuzz04(fizzAt, buzzAt int64) []string](<#func-fizzbuzz04>)
- [func FizzBuzzA(fizzAt, buzzAt int64) []string](<#func-fizzbuzza>)


## Constants

```go
const Buzz = "Y" // "Y"
```

const FizzBuzz = "Z" //"Z"

```go
const Digit = "O"
```

```go
const Fizz = "X" // "X"
```

## func FizzBuzz

```go
func FizzBuzz(fizzAt, buzzAt int64) []string
```

FizzBuzz performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\[\]

## func FizzBuzz01

```go
func FizzBuzz01(fizzAti, buzzAti int64) []string
```

FizzBuzz01 generates the expected array in an alternative way \(very inefficient\)\.

tag::fizz\-buzz\-01\[\]

## func FizzBuzz01A

```go
func FizzBuzz01A(fizzAti, buzzAti int64) []string
```

FizzBuzz01A generates the expected array in an alternative way \(very inefficient\)\.

tag::fizz\-buzz\-01\-a\[\]

## func FizzBuzz02

```go
func FizzBuzz02(fizzAt, buzzAt int64) []string
```

FizzBuzz02 performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\-02\[\]

## func FizzBuzz03

```go
func FizzBuzz03(fizzAt, buzzAt int64) []string
```

FizzBuzz03 performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\-03\[\]

## func FizzBuzz04

```go
func FizzBuzz04(fizzAt, buzzAt int64) []string
```

tag::fizz\-buzz\-04\[\]

## func FizzBuzzA

```go
func FizzBuzzA(fizzAt, buzzAt int64) []string
```

FizzBuzzA performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\-a\[\]




