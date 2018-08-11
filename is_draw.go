package main

func isDraw(board [][]string) bool {
	for _, row := range board {
		for _, chr := range row {
			if chr == "" {
				return false
			}
		}
	}
	return true
}
