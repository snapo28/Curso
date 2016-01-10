package main

import "fmt"

func main() {
	//declarar el array
	var palabras[15] string

	//mostrar el array
	fmt.Println(palabras)
	//mostrar la longitud del array
	fmt.Println(len(palabras))

	for i:=0; i<15; i++{
		palabras[i] = fmt.Sprint("Numero --> ",i)
	}

	//mostrar de nuevo el array
	fmt.Println(palabras)

	//ahora mostrarlas por separado
	for i:=0; i<len(palabras); i++{
		fmt.Println(palabras[i])
	}
}
