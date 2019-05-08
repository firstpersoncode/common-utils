package utils

import (
  "fmt"
  "math/rand"
)

func Random(min int, max int) int {
  return rand.Intn(max - min) + min
}

func Pad(value int) string {
  return fmt.Sprintf("%02d", value)
}
