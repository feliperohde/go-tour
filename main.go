package main

import (
  "fmt"
  "math/rand"
  "strings"
)

func main() {
  fmt.Println("my favorite number is", rand.Intn(10))
  fmt.Println(strings.Title("No, my favorite number is 4"))
}