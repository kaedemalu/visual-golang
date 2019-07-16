package main

import (
  // add packages
  "rand"
  "time"
  "math"
  "github.com/go-gl/glfw/v3.2/glfw"
)

const (
  BlockSize = 20
  FieldHeight = 20
  FieldWidth = 10
  TetroSize = 4
  WinWidth = BlockSize * FieldWidth
  WinHeight = BlockSize * FieldHeight
  TimerPeriod = 250 //ms
)

const (
  BTetros = [
          [66, 66, 66, 66],
          [27, 131, 72, 232],
          [36, 231, 36, 231],
          [63, 132, 63, 132],
          [311, 17, 223, 74],
          [322, 71, 113, 47],
          [1111, 9, 1111, 9],
  ]
  Colors = []
)

struct Block (
  x int
  y int
)

struct Game (
  pos_x int
  pos_y int
)

func main() {
  // add script
}

