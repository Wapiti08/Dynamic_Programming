package main

import (
	"fmt"
	"log"
)

/*
	define a Cache matrix to save the result

*/

func main() {
	var i, j int
	// input the N and M in the same line
	fmt.Scanln( &i, &j)
	// if (i<1 || float64(i)> math.Pow10(3)) || (j<1 || float64(j) > math.Pow10(3)){
	// 	fmt.Println("Please input the i and j again in range(1,10^3)")
	// 	return
	// }

	max_row := i
	max_column := j
	// create a matrix to save the input

	// take N and M as the range to input element in the matrix
	Matrix := make([][]int, max_row)
	var value int

	for row := range Matrix {
		Matrix[row] = make([]int, max_column)
		for column := range Matrix[row] {
			fmt.Scanf("%d",&value)
			Matrix[row][column] = value
		}
	}


	// input the number of queries
	var query int
	if _,err := fmt.Scan(&query); err != nil {
		log.Print("Scan for query failed, due to ", err)
		return
	}
	// if float64(query) > math.Pow10(5){
	// 	fmt.Println("Please input the number within 10^5")
	// 	return
	// }


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
