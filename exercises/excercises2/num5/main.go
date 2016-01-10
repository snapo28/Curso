package main

import (
	"fmt"
)

func main()  {
	m := []int {1,3,6,5,4,8,562,125,7,3}

	fmt.Println(variable(1,2,3,4))
	fmt.Println(variable())
	fmt.Println(variable(m...))
}

func variable(x ...int) (string){
	return "Inside variatic"
}