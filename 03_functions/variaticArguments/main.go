package main

import "fmt"

func main()  {
	//asignamos una lista de flotantes
	lista := []float64{15,35,0.1,3.21}

	//para pasar todos los datos de la lista se agrega tres puntos despues de el nombre
	x := promedio(lista...)
	fmt.Println("El resultado es: ", x)
}

func promedio(datos ...float64) (float64) {

	fmt.Println(datos)
	fmt.Printf("%T \n",datos)

	var total float64

	for _, v := range datos{
		total += v
	}
	return total / float64(len(datos))
}