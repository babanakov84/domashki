package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите количество строк:")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Введите количество столбцов:")
	fmt.Scanf("%d\n", &b)

	for i := 1; i <= a; i++ {
		switch i%2 == 1 {
		case true:
			for j := 1; j <= b; j++ {
				switch j%2 == 1 {
				case true:
					fmt.Printf("%3v", "X")
				default:
					fmt.Printf("%3v", " ")
				}
			}
			fmt.Printf("\n")
		case false:
			for j := 1; j <= b; j++ {
				switch j%2 == 1 {
				case true:
					fmt.Printf("%3v", " ")
				default:
					fmt.Printf("%3v", "X")
				}
			}
			fmt.Printf("\n")
		}
	}
}
