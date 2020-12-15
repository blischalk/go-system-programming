package anotherPackage

import (
  "fmt"
)

const version = "1.1"
const Pi = "3.14159"

func Add(x, y int) int {
  return x + y
}

func Println(x int) {
  fmt.Println(x)
}

func Version() {
  fmt.Println("The version of the package is", version)
}

func init() {
  fmt.Println("The init function of anotherPackage")
}
