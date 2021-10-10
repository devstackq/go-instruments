package handlers

import (
	"fmt"
)

const (
	PI = 3.14
)

type Foo interface {
	Bar() string
	Baz() string // new method added, which breaks the code
}

type Triangle struct {
	X, Y, Z int
 }

 type Point struct {
	R int
}

func DoThing(f Foo) {
	fmt.Println("bar call for each struct:", f.Bar(), f.Baz())
}

func NewTriangle()*Triangle {
	return &Triangle{}
}
// p := models.NewPoint()

func (t *Triangle) Bar() string {
	return fmt.Sprintf("triangle perim=%d", t.X + t.Y+ t.Z)
}

func (t *Triangle) Baz() string {
	return fmt.Sprintf("triangle area=%d", t.X * t.Y /2)
}


func (p *Point) Bar() string {
	return fmt.Sprintf("perim circle=%d", 2 * PI * float64(p.R))
}

func (p *Point) Baz() string {
	return fmt.Sprintf("area circle =%d", PI * float64(p.R))
}