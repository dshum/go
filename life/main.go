package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 200
	height = 50
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// Universe является двухмерным полем клеток.
type Universe [][]bool

// NewUniverse возвращает пустую вселенную.
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Seed заполняет вселенную случайными живыми клетками.
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

// Set устанавливает состояние конкретной клетки.
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// Alive сообщает, является ли клетка живой.
// Если координаты за пределами вселенной, возвращаемся к началу.
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// Neighbors подсчитывает прилегающие живые клетки.
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

// Next возвращает состояние определенной клетки на следующем шаге.
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// String возвращает вселенную как строку
func (u Universe) String() string {
	var b rune
	buf := make([]rune, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '☀'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}

	return string(buf)
}

// Show очищает экран и возвращает вселенную.
func (u Universe) Show() {
	fmt.Print("\033[H\033[2J", colorGreen, u.String(), colorReset)
}

// Step обновляет состояние следующей вселенной (b) из
// текущей вселенной (a).
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 100000; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 20)
		a, b = b, a // Swap universes
	}
}
