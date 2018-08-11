package main

import "testing"

func Test_Initial_Board_2Dimension_Size3_ShouldBeReturn_ArrayString2D_Size3(t *testing.T) {
	expected := [][]string{
		[]string{"", "", ""},
		[]string{"", "", ""},
		[]string{"", "", ""},
	}

	boardSize := 3

	_, actualBoard := InitialBoard(boardSize)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if expected[i][j] != actualBoard[i][j] {
				t.Errorf("expected %v but it got %v", expected[i][j], actualBoard[i][j])
			}
		}
	}
}

func Test_Initial_Board_2Dimension_Size2_ShouldBeReturn_ArrayString2D_Size2(t *testing.T) {
	expected := [][]string{
		[]string{"", ""},
		[]string{"", ""},
	}

	boardSize := 2

	_, actualBoard := InitialBoard(boardSize)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if expected[i][j] != actualBoard[i][j] {
				t.Errorf("expected %v but it got %v", expected[i][j], actualBoard[i][j])
			}
		}
	}

}

func Test_Initial_Board_2Dimension_Size1_ShouldBe_Error_LowOfSize(t *testing.T) {
	expected := "size of low"

	boardSize := 1

	err, _ := InitialBoard(boardSize)

	if expected != err.Error() {
		t.Errorf("expected %v but it got %v", expected, err.Error())
	}
}

func Test_Initial_Board_2Dimension_Size4_ShouldBe_Error_MoreOfSize(t *testing.T) {
	expected := "size of more than limit"

	boardSize := 4

	err, _ := InitialBoard(boardSize)

	if expected != err.Error() {
		t.Errorf("expected %v but it got %v", expected, err.Error())
	}
}
