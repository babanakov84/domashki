package main

import (
	"fmt"

	"github.com/babanakov84/domashki/hw06_testing/pkg/chessboard"
	"github.com/babanakov84/domashki/hw06_testing/pkg/fixapp"
	"github.com/babanakov84/domashki/hw06_testing/pkg/shapes"
	"github.com/babanakov84/domashki/hw06_testing/pkg/structcomparator"
)

func main() {
	// Примеры работы с пакетом shapes
	circle := shapes.Circle{Radius: 5}
	rectangle := shapes.Rectangle{Width: 10, Height: 5}
	triangle := shapes.Triangle{Base: 8, Height: 6}

	areaCircle, err := shapes.CalculateArea(circle)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Круг: радиус %.2f. Площадь: %.2f\n", circle.Radius, areaCircle)
	}

	areaRectangle, err := shapes.CalculateArea(rectangle)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Прямоугольник: ширина %.2f, высота %.2f. Площадь: %.2f\n",
			rectangle.Width, rectangle.Height, areaRectangle)
	}

	areaTriangle, err := shapes.CalculateArea(triangle)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Треугольник: основание %.2f, высота %.2f. Площадь: %.2f\n", triangle.Base, triangle.Height, areaTriangle)
	}

	// Примеры работы с пакетом chessboard
	var size int
	fmt.Print("Введите размер шахматной доски: ")
	_, err = fmt.Scanln(&size)
	if err != nil || size < 1 {
		fmt.Println("Введите число больше нуля!")
		return
	}

	err = chessboard.PrintBoard(size)
	if err != nil {
		fmt.Println(err)
	}

	// Пример работы с пакетом structcomparator

	book1 := structcomparator.Book{
		ID:     1,
		Title:  "Бойцовский клуб",
		Author: "Чак Паланик",
		Year:   2002,
		Size:   256,
		Rate:   7.8,
	}
	book2 := structcomparator.Book{
		ID:     2,
		Title:  "Капиталъ",
		Author: "Карл Маркс",
		Year:   1867,
		Size:   480,
		Rate:   8.57,
	}

	comparator := structcomparator.NewBookComparator(structcomparator.ByYear)
	fmt.Println("Сравнение по году:", comparator.Compare(&book1, &book2))

	comparator = structcomparator.NewBookComparator(structcomparator.BySize)
	fmt.Println("Сравнение по страницам:", comparator.Compare(&book1, &book2))

	comparator = structcomparator.NewBookComparator(structcomparator.ByRate)
	fmt.Println("Сравнение по рейтингу:", comparator.Compare(&book1, &book2))

	// Пример работы fixapp

	staff, err := fixapp.Fixapp("pkg/fixapp/data.json")
	if err != nil {
		fmt.Println(err)
	}
	_, err = fixapp.Fixapp("pkg/fixapp/baddata.json")

	fmt.Println(err)
	_, err = fixapp.Fixapp("pkg/fixapp/baddata.json")

	fmt.Println(err)
	fmt.Println(staff)
}
