package main

import "testing"

func Test_Three_IsSortBy_Row_SHOULDBE_Return_IsWinner_Is_True(t *testing.T) {
	expected := true

	board := [][]string{
		[]string{"x", "x", "x"},
		[]string{"y", "y", ""},
		[]string{"", "", ""},
	}

	actualOutput := isWinner(board)
	if expected != actualOutput {
		t.Errorf("expected %s but it got %s", expected, actualOutput)
	}
}

func Test_Three_IsNotSortBy_Row_SHOULDBE_Return_IsWinner_Is_False(t *testing.T) {
	expected := false

	board := [][]string{
		[]string{"x", "x", ""},
		[]string{"y", "y", ""},
		[]string{"", "", "x"},
	}

	actualOutput := isWinner(board)
	if expected != actualOutput {
		t.Errorf("expected %s but it got %s", expected, actualOutput)
	}
}

func Test_Three_IsNotSortBy_Column_SHOULDBE_Return_IsWinner_Is_True(t *testing.T) {
	expected := true

	board := [][]string{
		[]string{"x", "y", "y"},
		[]string{"x", "y", ""},
		[]string{"x", "", "x"},
	}

	actualOutput := isWinner(board)
	if expected != actualOutput {
		t.Errorf("expected %s but it got %s", expected, actualOutput)
	}
}

func Test_Three_IsNotSortBy_Diag_SHOULDBE_Return_IsWinner_Is_True(t *testing.T) {
	expected := true

	board := [][]string{
		[]string{"x", "y", "y"},
		[]string{"y", "x", ""},
		[]string{"x", "", "x"},
	}

	actualOutput := isWinner(board)
	if expected != actualOutput {
		t.Errorf("expected %s but it got %s", expected, actualOutput)
	}
}
