# go-parametrized-fizzbuzz
Concurrent Golang Solution for Parametrized Fizzbuzz. Unit Testing for All Possible Input Combinations: How to Build and Deploy Production-Grade Concurrent Go (Golang) projects.

View on Amazon
[Concurrent Golang Solution for Parametrized Fizzbuzz. Unit Testing for All Possible Input Combinations: How to Build and Deploy Production-Grade Concurrent Go (Golang) projects.](https://www.amazon.com/dp/B09ZTD459Z)
<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# fizzbuzz

```go
import "mobiledatabooks.com/go-parametrized-fizzbuzz/fizzbuzz"
```

## Index

- [Constants](<#constants>)
- [func FizzBuzz(fizzAt, buzzAt int64) []string](<#func-fizzbuzz>)
- [func FizzBuzzBigInt(fizzAt, buzzAt int64) []string](<#func-fizzbuzzbigint>)
- [func FizzBuzzConcurrent(fizzAt, buzzAt int64) []string](<#func-fizzbuzzconcurrent>)
- [func FizzBuzzContinue(fizzAt, buzzAt int64) []string](<#func-fizzbuzzcontinue>)
- [func FizzBuzzIfElse(fizzAt, buzzAt int64) []string](<#func-fizzbuzzifelse>)
- [func FizzBuzzMinMemory(fizzAt, buzzAt int64) []string](<#func-fizzbuzzminmemory>)
- [func FizzBuzzSlow(fizzAti, buzzAti int64) []string](<#func-fizzbuzzslow>)
- [func FizzBuzzSlowMinMemory(fizzAti, buzzAti int64) []string](<#func-fizzbuzzslowminmemory>)
- [func filter(c <-chan string, n int, label string) <-chan string](<#func-filter>)
- [func generate() <-chan string](<#func-generate>)


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

end::fizz\-buzz\-a\[\]

```go
const empty = "" //"x"
```

## func FizzBuzz

```go
func FizzBuzz(fizzAt, buzzAt int64) []string
```

FizzBuzz performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\[\]

## func FizzBuzzBigInt

```go
func FizzBuzzBigInt(fizzAt, buzzAt int64) []string
```

tag::fizz\-buzz\-04\[\]

## func FizzBuzzConcurrent

```go
func FizzBuzzConcurrent(fizzAt, buzzAt int64) []string
```

FizzBuzzConcurrent  Concurrent FizzBuzz with Go Routines from Russ Cox\. Here you find a parametrized solution\.

\.FizzBuzzConcurrent tag::fizz\-buzz\-FizzBuzzConcurrent\[\]

## func FizzBuzzContinue

```go
func FizzBuzzContinue(fizzAt, buzzAt int64) []string
```

FizzBuzzContinue performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\-03\[\]

## func FizzBuzzIfElse

```go
func FizzBuzzIfElse(fizzAt, buzzAt int64) []string
```

FizzBuzzIfElse performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\-02\[\]

## func FizzBuzzMinMemory

```go
func FizzBuzzMinMemory(fizzAt, buzzAt int64) []string
```

FizzBuzzMinMemory performs a FizzBuzz operation over a range of integers tag::fizz\-buzz\-a\[\]

## func FizzBuzzSlow

```go
func FizzBuzzSlow(fizzAti, buzzAti int64) []string
```

FizzBuzzSlow generates the expected array in an alternative way \(very inefficient\)\.

tag::fizz\-buzz\-01\[\]

## func FizzBuzzSlowMinMemory

```go
func FizzBuzzSlowMinMemory(fizzAti, buzzAti int64) []string
```

FizzBuzzSlowMinMemory generates the expected array in an alternative way \(very inefficient\)\.

tag::fizz\-buzz\-01\-a\[\]

## func filter

```go
func filter(c <-chan string, n int, label string) <-chan string
```

\<10\>

## func generate

```go
func generate() <-chan string
```


