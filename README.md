# go-parametrized-fizzbuzz
Golang Solution for Parametrized Fizzbuzz. Unit Testing for All Possible Input Combinations: How to Build and Deploy Production-Grade Go (Golang) projects.


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


