package main

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  var s [3]string
  var threshold = 0

  arguments := os.Args
  if len(arguments) == 2 {
    temp, err := strconv.Atoi(arguments[1])
    if err == nil {
      threshold = temp
    } else {
      log.Fatal(err)
    }
  }

  s[0] = "1 b 3 1 a a b"
  s[1] = "11 a 1 1 1 1 a a"
  s[2] = "-1 b 1 -4 a 1"

  counts := make(map[string]int)

  for i := 0; i < len(s); i++ {
    data := strings.Fields(s[i])
    for _, word := range data {
      _, ok := counts[word]
      if ok {
        counts[word] = counts[word] + 1
      } else {
        counts[word] = 1
      }
    }
  }

  for key, _ := range counts {
    if counts[key] >= threshold {
      fmt.Printf("%s -> %d \n", key, counts[key])
    }
  }
}
