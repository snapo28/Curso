package main

import "fmt"

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main ()  {
	fmt.Println("binario\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t",KB)
	fmt.Printf("%d\n",KB)
	fmt.Printf("%b\t\t",MB)
	fmt.Printf("%d\n",MB)


}
