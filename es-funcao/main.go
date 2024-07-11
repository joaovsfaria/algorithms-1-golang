package main

import "fmt"
import "math"

func main() {

var x float64

fmt.Println("Insira o valor de x")
fmt.Scan(&x)

y := math.Sqrt((1/(1+math.Pow(x,2))))
fmt.Printf("y = %.4f", y)

}
