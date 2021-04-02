/*
Given an integer N, print the factorial of the N (mod).

Input:
First line contains one integer, T, number of test cases.
Each test case contains one integer, N.

Output:
For each test case you need to print the factorial of N (mod
). 

*/

package main

import (
	"fmt"
	"log"
    "math"
)

func factorial(N uint64, Cache []uint64) uint64 {
	if Cache[N] != 0 {
		return Cache[N]
	}
	if N <= 1 {
		return 1
	}

	Cache[N] = uint64(math.Mod(float64(N * factorial(N-1, Cache)) , math.Pow10(9)+7))

	//fmt.Println(Cache)

	return Cache[N]
}

func main() {
	// input the first num for total inputs
	var i uint64

	if _, err := fmt.Scan(&i); err != nil {
		log.Print("  Scan for i failed, due to ", err)
		return
}

    Cache := make([]uint64, i+1)

    for ord := i; ord>=1; ord--{
        //fmt.Println("Current num is",ord)
        var input_num uint64
        if _, err := fmt.Scan(&input_num); err != nil {
            log.Print("  Scan for i failed, due to ", err)
        return
        } else {
            //fmt.Println(input_num)    
            fmt.Println(factorial(input_num, Cache))    
            }       
        }
}
