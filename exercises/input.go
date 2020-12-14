package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Provide me some input")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  line := scanner.Text()
  fmt.Printf("You entered: %q\n", line)
}
