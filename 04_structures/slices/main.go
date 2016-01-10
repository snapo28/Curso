package main

import "fmt"

func main() {
	//crear un slice con make
	//make(tipo,longitud,capacidad)
	//cuando no se asigna la capacidad automaticamente se le pone la de la longitud
	mySlice := make([]int,3)

	//asignar valores
	mySlice[0] = 2
	mySlice[1] = 9
	mySlice[2] = 7

	//mostrar los elementos
	fmt.Println(mySlice)

	//crear un slice con make
	languages := make([]string,3,10)
	languages[0] = "spanish"
	languages[1] = "english"
	languages[2] = "german"

	//mostrar de elemento por elemento
	fmt.Println(languages[0])
	fmt.Println(languages[1])
	fmt.Println(languages[2])

	//mostrar la capacidad
	//la capacidad aumenta automaticamente cuando excede el numero de elementos asiganos
	//esto a costo de operaciones porque crea un nuevo array de el doble de la capacidad anterior
	fmt.Println("la capacidad es:",cap(languages))

}
