/**
 * Name: Yu Heng Wu
 * VUnetID: wuy54
 * Email: yu.heng.wu@vanderbilt.edu
 * Class: CS 3270-02 - Vanderbilt University
 * Date: 04/18/2022
 * Honor statement: I have neither given nor received unauthorized aid.
 * Created by spenc on 4/18/2022.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

/**
* Read from a file and transform that into a sudoku board (represented by a 9x9 2d slice of int)
* Return true if it does read a file and successfully turn that into a sudoku board, false if not.
*
* @param board  A 9x9 2d slice of int that represents the sudoku board
* @return A bool value representing that if the sudoku board initialized successfully and should the program continue
 */
func readFile(board [][]int) bool {
	fileName := introInterface()

	if fileName != "quit" {
		fileName = absPath(fileName)

		var file *os.File

		fileName, file = fileOpened(fileName)

		if fileName != "quit" {
			setBoard(board, file)
			return true
		}

	}
	return false
}

/**
* Print the initial user interface to screen and ask an input value of the name of the file that user wants to open
*
* @return A string that is what the user input
 */
func introInterface() string {
	var fileName string
	fmt.Println("\nEnter the file's name: (file assumed in the project's txt folder)")
	fmt.Println("Pressing <Enter> will run the file \"sudoku-test1.txt\".")
	fmt.Println("Enter \"quit\" to terminate.")
	fmt.Scanln(&fileName)

	return fileName
}

/**
* Print out the sudoku board in a beautiful format.
*
* @param board  A 9x9 2d slice of int that represents the sudoku board
 */
func printBoard(board [][]int) {

	for i := 0; i < 9; i++ {
		if i != 0 && i%3 == 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < 8; j++ {
			if j != 0 && j%3 == 0 {
				fmt.Print("| ")
			}

			fmt.Print(board[i][j], " ")
		}
		fmt.Println(board[i][8])
	}

	fmt.Print("\n")
}

/**
* Turn the filepath(string) into absolute path. If fileName is an empty string, the path will be of the default file.
*
* @param fileName The name of the file the user want to open.
* @return An string that is the absolute path of the fileName
 */
func absPath(fileName string) string {
	if fileName == "" {
		fileName, _ = filepath.Abs("txt/sudoku-test1.txt")
	} else {
		fileName, _ = filepath.Abs("txt/" + fileName)
	}

	return fileName
}

/**
* Open the file and return the opened file if opened successfully
*
* @param fileName The name of the file the user want to open.
* @return An opened reference of os.File
 */
func fileOpened(fileName string) (string, *os.File) {
	file, err := os.Open(fileName)
	for err != nil {
		fileName = ""
		fmt.Println("\nFailed to open the file. \nRe-enter the file's name: ")
		fmt.Scanln(&fileName)

		if fileName == "quit" {
			break
		}

		fileName = absPath(fileName)

		file, err = os.Open(fileName)
	}

	return fileName, file
}

/**
* Set the sudoku board according to the .txt opened that passed int
*
* @param board  A 9x9 2d slice of int that represents the sudoku board
* @param file A reference of an opened os.File
 */
func setBoard(board [][]int, file *os.File) {
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			scanner.Scan()

			if len(scanner.Text()) > 1 {
				log.Panicf("Input's not a valid Sudoku Board")
			} else if num, err := strconv.Atoi(scanner.Text()); err == nil {
				board[i][j] = num
			} else {
				log.Panicf("Input's not a valid Sudoku Board")
			}

		}
	}
}
