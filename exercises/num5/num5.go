package main

import "fmt"

func main()  {
	var suma int
	for i:=0;i<1000;i++{
		if i%3 ==0 || i%5 ==0{
			fmt.Println(i)
			suma += i
		}
	}
	fmt.Println("La suma de los multiplos es: ",suma)
}
