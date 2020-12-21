package main
import (
  "flag"
  "fmt"
  "os"
)

func main() {
  minusVERBOSE := flag.Bool("v", false, "Verbose")
  flag.Parse()
  arguments := flag.Args()

  if len(arguments) == 0 {
    fmt.Println("Please provide an argument!")
    os.Exit(1)
  }

  for _, file := range arguments {

    if *minusVERBOSE {
      fmt.Println("INFO: Preparing to remove:", file)
    }

    err := os.Remove(file)
    if err != nil {
      if *minusVERBOSE {
        fmt.Println("ERROR: Failed to remove:", file)
      }
      fmt.Println(err)
      continue
    } else {
      if *minusVERBOSE {
        fmt.Println("INFO: Success! Removed:", file)
      }
    }
  }
}
