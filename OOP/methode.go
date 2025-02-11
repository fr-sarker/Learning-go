package main

import "fmt"

type NameObj struct {
	Name string
}

func (n NameObj) show() {
	fmt.Println(n.Name)
}

type Rectangle struct {
	NameObj
	Width, Height float64
}

func (r Rectangle) show() {
	fmt.Println("Rectangle", r.Name, r.NameObj.Name)
}

func main() {
	var a = NameObj{"Joe"}
	var b = Rectangle{NameObj{"Richard"}, 10, 20}
	a.show()
	b.show()
}
