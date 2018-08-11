package main

import "testing"

func Test_NotNull_In_Array_ShouldBe_Return_IsDraw_True(t *testing.T) {
	expected := true

	board := [][]string{
		[]string{"x", "y", "y"},
		[]string{"y", "x", "x"},
		[]string{"x", "y", "y"},
	}

	actualOutput := isDraw(board)

	if expected != actualOutput {
		t.Errorf("expected %s but it got %s", expected, actualOutput)
	}
}

func Test_Null_In_Array_ShouldBe_Return_IsDraw_false(t *testing.T) {
	expected := false

	board := [][]string{
		[]string{"x", "y", "y"},
		[]string{"y", "x", "x"},
		[]string{"x", "", "y"},
	}

	actualOutput := isDraw(board)

	if expected != actualOutput {
		t.Errorf("expected %s but it got %s", expected, actualOutput)
	}
}
