package designpattern

import "fmt"

// Lets learn facotry design pattern with below example

// here we are going to take example of create a shaep based on type passed into the create shape function

// lets have one interface

type Shape interface {
	Draw()
}

type Reactangle struct {
	Height int
	Width  int
}

func (r *Reactangle) Draw() {
	result := r.Height * r.Width
	fmt.Println("Drawing a Rectangle with area:", result)
}

type Square struct {
	Side float64
}

func (s Square) Draw() {
	fmt.Println("Drawing a square with side:", s.Side)
}

type ShapeFactory struct{}

func (sh ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "Reactangle":
		return &Reactangle{Height: 5, Width: 6}
	case "Square":
		return &Square{Side: 5}
	default:
		return nil
	}
}
