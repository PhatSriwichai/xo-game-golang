package main

import (
	"fmt"
)

type XOGame struct {
	xPlayer       Player
	yPlayer       Player
	board         [][]string
	isTurnPlayerX bool
	dimension     int
	winner        Player
	status        string
}

func (xoGame *XOGame) InitialBoard(dimension int) error {
	err, board := InitialBoard(dimension)
	if err != nil {
		fmt.Errorf("Initial game error: %s\n", err.Error())
		return err
	}
	xoGame.board = board
	xoGame.status = "start"
	return nil
}

func (xoGame *XOGame) switchPlayer() {
	xoGame.isTurnPlayerX = !xoGame.isTurnPlayerX
}

func (xoGame *XOGame) isBoardPositionEmpty(xPosition int, yPosition int) bool {
	if xoGame.board[xPosition][yPosition] == "" {
		return true
	}
	return false
}

func (xoGame *XOGame) move(xPosition int, yPosition int) string {
	var char string = "x"
	if !xoGame.isTurnPlayerX {
		char = "y"
	}
	xoGame.board[xPosition][yPosition] = char
	return xoGame.board[xPosition][yPosition]
}

func (xoGame *XOGame) isWinner() {
	isWin := isWinner(xoGame.board)
	if isWin {
		xoGame.status = "end"
		if xoGame.isTurnPlayerX {
			xoGame.winner = xoGame.xPlayer
		} else {
			xoGame.winner = xoGame.yPlayer
		}
	}
}

func (xoGame *XOGame) isDraw() {
	if isDraw(xoGame.board) {
		xoGame.status = "end"
	}
}
