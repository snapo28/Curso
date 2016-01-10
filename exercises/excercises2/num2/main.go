package main

import "fmt"

func main()  {
	//function expression
	twoReturns := func (x int) (float64,bool){
		return float64(x)/2, x%2 == 0
	}
	fmt.Println(twoReturns(5))
}
