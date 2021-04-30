package main

import (
	"fmt"
)

func main() {
	var max_row, max_column int
	// input the N and M in the same line
	fmt.Scanln( &max_row, &max_column)

	// take N and M as the range to input element in the matrix
	var Matrix [1001][1001]int
	var value int

	for x:=0; x<max_row; x++{
		for y:=0; y<max_column;y++ {
			fmt.Scanf("%d",&value)
			Matrix[x][y] = value
		}
	}

	// input the number of queries
	var query int
	fmt.Scanf("%d",&query)
		
	// creat the Cache Matrix to save computed results
	for column:=1; column < max_column; column++ {
		Matrix[0][column] += Matrix[0][column-1]
	}
	
	for row:=1; row < max_row; row++ {
		Matrix[row][0] += Matrix[row-1][0]
	}

	for row:=1; row<max_row; row++ {
		for column:=1;column<max_column;column++{
				Matrix[row][column] += Matrix[row-1][column] + Matrix[row][column-1]-
				Matrix[row-1][column-1]
			}
		}
	// input the location depending on the times of queries
	for ord := query; ord>=1; ord -- {
		var num1, num2 int
		// input the query nums
		fmt.Scanf("%d %d",&num1, &num2)
		fmt.Println(Matrix[num1-1][num2-1])
	}
}
