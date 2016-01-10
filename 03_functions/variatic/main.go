package main

import "fmt"

func main()  {
	x := promedio(15,35,0.1,3.21)
	fmt.Println("El resultado es: ", x)
}
//para recibir un numero infinito de argumentos se asignan tres puntos antes del parametro
func promedio(datos ...float64) (float64) {

	fmt.Println(datos)
	fmt.Printf("%T \n",datos)

	var total float64

	for _, v := range datos{
		total += v
	}
	return total / float64(len(datos))
}