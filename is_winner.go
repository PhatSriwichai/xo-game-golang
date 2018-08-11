package main

func isWinner(board [][]string) bool {
	for _, row := range board {
		if rowVictoryCheck(row) {
			return true
		}
	}

	for i := 0; i < len(board); i++ {
		column := boardColumn(board, i)
		if rowVictoryCheck(column) {
			return true
		}
	}

	return diagVitoryCheck(board)
}

func diagVitoryCheck(board [][]string) bool {
	dimension := len(board)
	diag := make([]string, dimension)
	for i := 0; i < dimension; i++ {
		diag[i] = board[i][i]
	}

	if rowVictoryCheck(diag) {
		return true
	}

	diag = make([]string, dimension)
	for i := 0; i < dimension; i++ {
		diag[i] = board[i][dimension-i-1]
	}

	if rowVictoryCheck(diag) {
		return true
	}
	return false
}

func boardColumn(board [][]string, columnIndex int) (column []string) {
	column = make([]string, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return column
}

func rowVictoryCheck(row []string) bool {
	firstChr := ""
	for i, chr := range row {
		if chr == "" {
			return false
		}
		if i == 0 {
			firstChr = chr
			continue
		}

		if chr != firstChr {
			return false
		}
	}
	return true
}
