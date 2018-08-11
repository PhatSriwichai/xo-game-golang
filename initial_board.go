package main

import "errors"

func InitialBoard(size int) (error, [][]string) {
	if size < 2 {
		return errors.New("size of low"), nil
	} else if size > 3 {
		return errors.New("size of more than limit"), nil
	}
	boards := make([][]string, size)
	for i := range boards {
		boards[i] = make([]string, size)
	}
	return nil, boards
}
