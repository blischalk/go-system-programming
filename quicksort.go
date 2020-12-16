package main

import (
  "fmt"
  "math/rand"
  "time"
)

func quickSort(a []int) []int {
  if len(a) < 2 {
    return a
  }

  left, right := 0, len(a) - 1
  pivot := rand.Int() % len(a)

  a[pivot], a[right] = a[right], a[pivot]

  for i, _ := range a {
    if a[i] < a[right] {
      a[left], a[i] = a[i], a[left]
      left++
    }
  }

  a[left], a[right] = a[right], a[left]

  quicksort(a[:left])
  quicksort(a[left+1])

  return a
}

func main() {
  fmt.Println("Starting quick sort program...");

  COLLECTION_SIZE := 100
  ELEMENT_RANGE := 100000
  elements := make([]int, COLLECTION_SIZE)

  rand.Seed(time.Now().Unix())
  for i := 0; i < COLLECTION_SIZE; i++ {
    myrand := rand.Intn(ELEMENT_RANGE)
    fmt.Printf("Adding: %d\n", myrand)
    elements[i] = myrand
  }

  for i := range elements {
    fmt.Printf("Element: %d is %d\n", i, elements[i])
  }

  fmt.Println("Sorting...");
  quickSort(elements)

  for i := range elements {
    fmt.Printf("Element: %d is %d\n", i, elements[i])
  }
}
