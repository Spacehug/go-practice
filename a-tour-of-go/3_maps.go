package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) (counter map[string]int) {
  counter = make(map[string]int)
  for _, word := range strings.Fields(s) {
    counter[word]++
  }
  return
}

func main() {
  wc.Test(WordCount)
}
