package main

import "fmt"

type base struct {
	a string
	b int
}

func (this base) xyz() {
	fmt.Println("xyz, a is:", this.a)
}

func (this base) display() {
	fmt.Println("base, a is:", this.a)
}

type derived struct {
	base // embedding
	d    int
	a    float32 // shadowed
}

// method disply
func (this derived) display() {
	fmt.Println("derived, a is:", this.a)
}

func main() {
	var a derived = derived{base{"a", 10}, 10, 10.5}
	a.display()
	a.base.display()
	a.xyz()
}
