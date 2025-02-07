package main

import "fmt"

type Circle struct {
	x, y, r float64
}

type Android struct {
	Circle
	Model string
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{10, 4, 5}
	d := new(Android)
	d.Circle = c

	fmt.Println(totalArea(&c, &d))
}
