package main

import "fmt"

func main() {

var raio float64

pi := 3.14159

fmt.Println("Insira o valor do raio")

fmt.Scan(&raio)

area := pi * raio * raio

fmt.Printf("A area do circulo é %v", area)

}
