[![Tests on Linux, MacOS and Windows](https://github.com/bep/predicate/workflows/Test/badge.svg)](https://github.com/bep/predicate/actions?query=workflow:Test)
[![Go Report Card](https://goreportcard.com/badge/github.com/bep/predicate)](https://goreportcard.com/report/github.com/bep/predicate)
[![GoDoc](https://godoc.org/github.com/bep/predicate?status.svg)](https://godoc.org/github.com/bep/predicate)


```go
var (
   pHello predicate.P[string] = func(s string) bool {
      return s == "hello"
   }
   pWorld predicate.P[string] = func(s string) bool {
      return s == "world"
   }
   pAny predicate.P[string] = func(s string) bool {
      return s != ""
   }
)

fmt.Println("Or (true):", pHello.Or(pWorld)("hello"))
fmt.Println("Or (false):", pHello.Or(pWorld)("foo"))
fmt.Println("And (false):", pHello.And(pWorld)("hello"))
fmt.Println("And (true):", pHello.And(pAny)("hello"))
fmt.Println("Negate (false):", pHello.Negate()("hello"))
fmt.Println("Negate (true):", pHello.Negate()("world"))
fmt.Println("Chained (true):", pHello.And(pAny.Or(pWorld))("hello"))
fmt.Println("Chained (false):", pHello.And(pAny.Or(pWorld))("foo"))
```