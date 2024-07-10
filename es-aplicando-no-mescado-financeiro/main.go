package main

import "fmt"
import "math"

func main() {


var c float64
var i float64
var t float64

fmt.Println("Insira o valor do capital, taxa de juros e do tempo em meses")

fmt.Scan(&c)
fmt.Scan(&i)
fmt.Scan(&t)

m := c* math.Pow((1 + i/100), t)


lucro := m - c

fmt.Printf("O lucro foi de R$%.2f", lucro)
}
