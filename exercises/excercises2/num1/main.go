package main

import "fmt"

func main()  {
	fmt.Println(twoReturns(5))
}

func twoReturns(x int) (float64,bool){
	return float64(x)/2, x%2 == 0
	}