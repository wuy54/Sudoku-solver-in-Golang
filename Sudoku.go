/**
 * Name: Yu Heng Wu
 * VUnetID: wuy54
 * Email: yu.heng.wu@vanderbilt.edu
 * Class: CS 3270-02 - Vanderbilt University
 * Date: 04/18/2022
 * Honor statement: I have neither given nor received unauthorized aid.
 * Created by spenc on 4/16/2022.
 */

package main

import (
	"fmt"
)

func main() {

	board := make([][]int, 9)

	for i := range board {
		board[i] = make([]int, 9)
	}

	if readFile(board) {
		fmt.Print("\nPuzzle: \n\n")

		printBoard(board)

		fmt.Print("\nSolution: \n\n")

		if !solveNum(0, 0, board) {
			fmt.Println("This sudoku puzzle is not solvable.")
		} else {
			printBoard(board)
		}
	}

	fmt.Println("\nProgram terminated.")
}

/**
* Solve the sudoku board
*
* @param row An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param col An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param board  A 9x9 2d slice of int that represents the sudoku board
* @return A bool value representing that if the sudoku board is solvable or not
 */
func solveNum(row uint, col uint, board [][]int) bool {

	numPresented := make([]bool, 10)

	if row > 8 {
		return true
	} else if col > 8 {
		if solveNum(row+1, 0, board) {
			return true
		}
	} else if board[row][col] == 0 {
		setNum(row, col, board, numPresented)

		if board[row][col] == 0 {
			return false
		} else if solveNum(row, col+1, board) {
			return true
		} else {
			for board[row][col] < 9 {
				numNow := board[row][col]
				numPresented[numNow] = true
				setNum(row, col, board, numPresented)

				if board[row][col] > numNow {
					if solveNum(row, col+1, board) {
						return true
					}
				} else {
					board[row][col] = 0
					return false
				}
			}
			board[row][col] = 0
		}
	} else {
		if solveNum(row, col+1, board) {
			return true
		}
	}

	return false
}

/**
* Set the specified cell of the sudoku board to a possible solution number
*
* @param row An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param col An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param board  A 9x9 2d slice of int that represents the sudoku board
 */
func setNum(row uint, col uint, board [][]int, numPresented []bool) {
	numRowCol(row, col, numPresented, board)
	numBlk(row, col, numPresented, board)

	idx := 1
	for idx < 10 && numPresented[idx] {
		idx++
	}

	if idx < 10 {
		board[row][col] = idx
	}
}

/**
* Check what numbers are already on the sudoku board of the row and column
* of the sudoku board's cell that we are trying to fill into.
* The numbers already on the board will be marked true in the numPresented (a bool slice)
*
* @param row An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param col An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param numPresented A bool slice that stores true if the number of an index is already on the board and is in the same
* row, column, or 3x3 block with the cell that we are trying to fill in.
* @param board  A 9x9 2d slice of int that represents the sudoku board
 */
func numRowCol(row uint, col uint, numPresented []bool, board [][]int) {
	for i := 0; i < 9; i++ {
		if board[i][col] > 0 {
			numPresented[board[i][col]] = true
		}
	}

	for j := 0; j < 9; j++ {
		if board[row][j] > 0 {
			numPresented[board[row][j]] = true
		}
	}
}

/**
* Check what numbers are already on the sudoku board of the 3x3 block
* of the sudoku board's cell that we are trying to fill into.
* The numbers already on the board will be marked true in the numPresented (a bool slice)
*
* @param row An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param col An uint indicating the row of the sudoku board's cell that the program is trying to fill in.
* @param numPresented A bool slice that stores true if the number of an index is already on the board and is in the same
* row, column, or 3x3 block with the cell that we are trying to fill in.
* @param board  A 9x9 2d slice of int that represents the sudoku board
 */
func numBlk(row uint, col uint, numPresented []bool, board [][]int) {

	rowDiff := row % 3
	colDiff := col % 3
	row -= rowDiff
	col -= colDiff

	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if board[i][j] > 0 {
				numPresented[board[i][j]] = true
			}
		}
	}
}
