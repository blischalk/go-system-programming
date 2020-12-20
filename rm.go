package main
import (
  "fmt"
  "os"
)

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("Please provide an argument!")
    os.Exit(1)
  }

  for idx, file := range arguments {

    if idx == 0 {
      continue
    }
    err := os.Remove(file)
    if err != nil {
      fmt.Println(err)
      continue
    }
  }
}
