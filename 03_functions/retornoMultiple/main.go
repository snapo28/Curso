package main

import "fmt"

func main()  {
	fmt.Println(saludo("Carlos ","Muñoz "))
}

func saludo(nombre string, apellido string) (string, string) {

	//Sprint lo que hace es concatenar
	return fmt.Sprint("Hola ",nombre,apellido), fmt.Sprint("\nHola ",apellido,nombre)
}