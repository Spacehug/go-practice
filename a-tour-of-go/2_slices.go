package main

import (
  "math"
  "golang.org/x/tour/pic"
)

func Pic(pattern func(x, y int) uint8) func(dx int, dy int) [][]uint8 {
  picture := func(dx int, dy int) [][]uint8 {
    pixels := make([][]uint8, dy)
    for y := range pixels {
      pixels[y] = make([]uint8, dx)
      for x := range pixels[y] {
        pixels[y][x] = pattern(x, y)
      }
    }
    return pixels
  }
  return picture
}

func a(x, y int) uint8 {
  return uint8(x * y)
}

func b(x, y int) uint8 {
  return uint8((x + y) / 2)
}

func c(x, y int) uint8 {
  return uint8(x ^ y)
}

func d(x, y int) uint8 {
  return uint8((x ^ y) * (x ^ y))
}

func e(x, y int) uint8 {
  return uint8((x ^ y) % 150)
}

func f(x, y int) uint8 {
  return (uint8(x ^ y) << uint8(y ^ x) - 16)
}

func g(x, y int) uint8 {
  return uint8(y - int(32 * math.Sin(float64(x * 2))) - 128)
}

func h(x, y int) uint8 {
  return uint8(((1 / 2) * y - x) ^ y - x)
}

func i(x, y int) uint8 {
  return uint8((y - x) * (y + x))
}

func j(x, y int) uint8 {
  return uint8(x * x + y * y)
}

func k(x, y int) uint8 {
  return uint8(x * x ^ y * y)
}

func l(x, y int) uint8 {
  return uint8(y ^ int(4 * math.Tan(float64(x ^ y))))
}

func main() {
  pic.Show(Pic(l))
}
