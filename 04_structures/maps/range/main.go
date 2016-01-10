package main

import "fmt"

func main() {
	//declarar el mapa
	myMapa := map[int] string{
		0: "Language",
		1: "Age",
		2: "Country",
		3: "Companie",
	}

	for key, val := range myMapa{
		fmt.Println(key," --> ",val)
	}
}
