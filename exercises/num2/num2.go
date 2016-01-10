package main

import "fmt"

func main()  {
	var small int
	var large int

	fmt.Print("Introduce un numero pequeño: ")
	fmt.Scan(&small)
	fmt.Print("Introduce un numero Mayor que el anterior: ")
	fmt.Scan(&large)

	fmt.Println(large, " / ", small, " = " , large/small)
}
