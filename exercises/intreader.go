package main


// Keep reading and storing integers from input until 0
// Print the min and max

import (
  "os"
  "fmt"
)

func main() {
  fmt.Println("Initializing int storage")
  var input int = -1
  theSlice := []int{}

  for true {
    fmt.Print("Please enter an integer: ")
    fmt.Scan(&input)
    if input != 0 {
      theSlice = append(theSlice, input)
      fmt.Println("Thank you. May I have another?")
    } else {
      break
    }
  }

  if len(theSlice) == 0 {
    os.Exit(0)
  }

  var MIN int = theSlice[0]
  var MAX int = theSlice[0]

  for _, number := range theSlice {
    if number > MAX {
      MAX = number
    }

    if number < MIN {
      MIN = number
    }
  }

  fmt.Printf("The MIN is: %d", MIN)
  fmt.Printf("The MAX is: %d", MAX)
}
