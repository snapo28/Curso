package main

import "fmt"

func main()  {
	fmt.Println(saludo("Carlos ","Muñoz"))
}

func saludo(nombre string, apellido string) string {
	return fmt.Sprint("Hola ",nombre,apellido)
}