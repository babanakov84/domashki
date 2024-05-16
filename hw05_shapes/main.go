package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() (float32, error)
}
type Circle struct {
	radius float32
}

func (c Circle) Area() (float32, error) {
	if c.radius < 0 {
		return 0, fmt.Errorf("ошибка: радиус не может быть отрицательный")
	}
	return math.Pi * c.radius * c.radius, nil
}

type Rectangle struct {
	width, height float32
}

func (r Rectangle) Area() (float32, error) {
	if r.width < 0 || r.height < 0 {
		return 0, fmt.Errorf("ошибка: ширина и высота не могут быть отрицательными")
	}
	return r.width * r.height, nil
}

type Triangle struct {
	base, height float32
}

func (t Triangle) Area() (float32, error) {
	if t.base < 0 || t.height < 0 {
		return 0, fmt.Errorf("ошибка: основание и высота не могут быть отрицательными")
	}
	return 0.5 * t.base * t.height, nil
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
	Circle1 := Circle{radius: 20}
	areaCircle, err := calculateArea(Circle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Круг: радиус %f \nПлощадь: %f", Circle1.radius, areaCircle)
	}
	Rectangle1 := Rectangle{width: 15, height: 12}
	areaRectangle, err := calculateArea(Rectangle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nПрямоугольник: ширина %f, высота %f \nПлощадь: %f", Rectangle1.width, Rectangle1.height, areaRectangle)
	}
	Triangle1 := Triangle{base: 14, height: 7}
	areaTriangle, err := calculateArea(Triangle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nТреугольник: основание %f, высота %f \nПлощадь: %f", Triangle1.base, Triangle1.height, areaTriangle)
	}
	Parallelogram1 := "parallelogram"
	_, err = calculateArea(Parallelogram1)
	if err != nil {
		fmt.Println("\nПараллелограмм:\n", err)
	}
}
