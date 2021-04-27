package main

import (
	"fmt"
	"log"
	"math"
)

/*
	define a Cache matrix to save the result

*/

func calculate_matrix(Matrix [][]int, Cache [][]int, num1 int, num2 int) int {

	if Cache[num1][num2]!= 0 {
		return Cache[num1][num2]
	}

	for row1 :=0; row1<=num1; row1++ {
		for column1 :=0; column1<=num2; column1++ {

			Cache[num1][num2] += Matrix[row1][column1]
		}
	}

	return Cache[num1][num2]
}

func main() {
	var i, j int
	// input the N and M in the same line
	fmt.Scanln( &i, &j)
	if (i<1 || float64(i)> math.Pow10(3)) || (j<1 || float64(j) > math.Pow10(3)){
		fmt.Println("Please input the i and j again in range(1,10^3)")
		return
	}

	max_row := i
	max_column := j
	// create a matrix to save the input
	
	// take N and M as the range to input element in the matrix
	Matrix := make([][]int, i)
	var value int
	cal := 1
	for row := range Matrix {
		Matrix[row] = make([]int, j)
		for column := range Matrix[row] {
			cal += 1
			fmt.Scanf("%d",&value)
			Matrix[row][column] = value
			if cal > 1000:
				
		}
	}


	// input the number of queries
	var query int
	if _,err := fmt.Scan(&query); err != nil {
		log.Print("Scan for query failed, due to ", err)
		return
	}
	if float64(query) > math.Pow10(5){
		fmt.Println("Please input the number within 10^5")
		return
	}


	// creat the Cache Matrix to save computed results
	Cache := make([][]int, max_row)
	for ord := range Matrix {
		Cache[ord] = make([]int, max_column)
	}
	// input the location depending on the times of queries
	for ord := query; ord>=1; ord -- {
		var num1, num2 int
		// input the query nums
		fmt.Scanf("%d %d",&num1, &num2)
		if (num1<1 || num1>i) || (num2<1 || num2>j) {
			fmt.Println("Please input the num1 and num2 again in range(i) and range(j)")
			return
	}
		fmt.Println(calculate_matrix(Matrix, Cache, num1-1, num2-1))
	}
}
