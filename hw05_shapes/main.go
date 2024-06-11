package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() (float32, error)
}
type Circle struct {
	Radius float32
}

func (c Circle) Area() (float32, error) {
	if c.Radius < 0 {
		return 0, fmt.Errorf("ошибка: радиус не может быть отрицательный")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

type Rectangle struct {
	Width, Height float32
}

func (r Rectangle) Area() (float32, error) {
	if r.Width < 0 || r.Height < 0 {
		return 0, fmt.Errorf("ошибка: ширина и высота не могут быть отрицательными")
	}
	return r.Width * r.Height, nil
}

type Triangle struct {
	Base, Height float32
}

func (t Triangle) Area() (float32, error) {
	if t.Base < 0 || t.Height < 0 {
		return 0, fmt.Errorf("ошибка: основание и высота не могут быть отрицательными")
	}
	return 0.5 * t.Base * t.Height, nil
}

func calculateArea(s any) (float32, error) {
	switch t := s.(type) {
	case Shape:
		return t.Area()
	default:
		return 0, fmt.Errorf("ОШИБКА!!!: переданный объект не является заданной фигурой")
	}
}

func main() {
	Circle1 := Circle{Radius: 20}
	areaCircle, err := calculateArea(Circle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Круг: радиус %f \nПлощадь: %f", Circle1.Radius, areaCircle)
	}
	Rectangle1 := Rectangle{Width: 15, Height: 12}
	areaRectangle, err := calculateArea(Rectangle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nПрямоугольник: ширина %f, высота %f \nПлощадь: %f", Rectangle1.Width, Rectangle1.Height, areaRectangle)
	}
	Triangle1 := Triangle{Base: 14, Height: 7}
	areaTriangle, err := calculateArea(Triangle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nТреугольник: основание %f, высота %f \nПлощадь: %f", Triangle1.Base, Triangle1.Height, areaTriangle)
	}
	Parallelogram1 := "parallelogram"
	_, err = calculateArea(Parallelogram1)
	if err != nil {
		fmt.Println("\nПараллелограмм:\n", err)
	}
}
