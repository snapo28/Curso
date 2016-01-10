package main

import "fmt"

func main()  {
	saludo("Carlos","Muñoz")
}

func saludo(nombre string, apellido string)  {
	fmt.Println("Hola",nombre,apellido)
}