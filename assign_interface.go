/*
1. Create two custom struct types called 'triangle' and 'square'
2. Square struct: a filed called 'sideLength' of type float64
3. Trangle struct: a filed called 'height' of tyoe float64 and filed of type 'base' of type float64
3. Both type should have function called 'getArea' returns the area
	trangle area: 0.5 * base * height
	square area: sidelength * sidelength
4. Add a 'shape' interface that defines a function called 'getArea'
5. printArea() to print out the area amd ot can be called either a triangle or a square.
*/
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t) // 50
	printArea(s) // 100
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
