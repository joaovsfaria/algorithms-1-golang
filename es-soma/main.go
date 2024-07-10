package main

import "fmt"

func main() {

var a, b int

fmt.Println("Insira o valor de 'a' e 'b'")

fmt.Scan(&a)
fmt.Scan(&b)

soma := a + b

fmt.Printf("O a soma dos valores é igual a %v", soma)

}
