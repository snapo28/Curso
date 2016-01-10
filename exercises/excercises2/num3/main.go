package main

import (
	"fmt"
)

func main()  {
	m := []int {1,3,6,5,4,8,562,125,7,3}

	fmt.Println("El numero mas grande es: ",maximo(m...))
}

func maximo(x ...int) (int){
	var numFinal int
	for _, v := range x {
		if v > numFinal{
			numFinal = v
		}
	}
	return numFinal
	}