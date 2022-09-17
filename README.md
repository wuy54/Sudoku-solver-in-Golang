# Sudoku-solver-in-Golang
_________________________
This is a sudoku solver. <br />
  To use this in a shell. After compiling and executing the program, you'll need to enter the location of a .txt file of which contains the sudoku you want it to solve. <br />
  <br />
  For the format of the sudoku board, please check the example .txt files in the txt folder.

<br />

## Instruction
### 1. How to start the program
1. For running and testing the program, please first download and install GO at https://go.dev
2. To check the success of GO's installation, type the below in your terminal
```
go version
```
3. Git clone the whole project to your computer
4. Open the terminal for your operating system
5. Go to the project folder in the terminal
6. To make the program run, type the below in your terminal
```
go run Sudoku.go
```
### 2. How to use the solver
## Instruction
### 1. How to start the program
1. For running and testing the program, please first download and install GO at https://go.dev
2. To check the success of GO's installation, type the below in your terminal
```
go version
```
3. Git clone the whole project to your computer
4. Open the terminal for your operating system
5. Go to the project folder in the terminal
6. To make the program run, type the below in your terminal
```
go run .
```
### 2. How to use the solver
* Please put the sudoku you want to solve in .txt file in the txt folder of the project
* The txt file must be in the format below. The 0s represent empty cells
```
0 4 3 0 8 0 2 5 0
6 0 0 0 0 0 0 0 0
0 0 0 0 0 1 0 9 4
9 0 0 0 0 4 0 7 0
0 0 0 6 0 8 0 0 0
0 1 0 2 0 0 0 0 3
8 2 0 5 0 0 0 0 0
0 0 0 0 0 0 0 0 5
0 3 4 0 9 0 7 1 0
```
* When the program is running successfully in your terminal, the program will show the below message:
```
Enter the file's name: (file assumed in the project's txt folder)
Pressing <Enter> will run the file "sudoku-test1.txt".
Enter "quit" to terminate.
```
* As the message showed, press \<Enter> for running the default file "sudoku-test1.txt" or enter the file's name in the same format
* Enter exactly "quit" to immediately terminate the program.
* If accidentally entered the non-existing file name or removed the file from txt folder, it will show the messages:
```
Failed to open the file.
Re-enter the file's name:
```
* The commands you can use is the same: 
  1. \<Enter> for default test case 
  2. Type the specific file name
  3. "quit" to terminate the program

