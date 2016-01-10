package main

import "fmt"

func main()  {
	x := make([]string, 1, 25)
	fmt.Println("array vacio:",x)
	cambiar(x)
	fmt.Println("array dentro del main:",x)
}

func cambiar(z []string){
	z[0] = "Carlos"
	fmt.Println("array dentro de funcion:",z)
}