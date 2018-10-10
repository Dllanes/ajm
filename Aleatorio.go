package ajm

import (
  "fmt"
)

func LastElement ( arr []int) int {
  return arr[len(arr)-1]
}

func PrintArray (arr []int) {
  fmt.Println(arr)
}
