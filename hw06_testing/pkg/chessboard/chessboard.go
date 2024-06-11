package chessboard

import (
	"errors"
	"fmt"
)

func CreateBoard(size int) (string, error) {
	if size <= 0 {
		return "", errors.New("размер доски должен быть больше нуля")
	}
	board := ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += " "
			} else {
				board += "#"
			}
		}
		board += "\n"
	}
	return board, nil
}

func PrintBoard(size int) error {
	board, err := CreateBoard(size)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(board)
	return nil
}
