package main

import "fmt"
import "math"

func main() {

var x, y float64

fmt.Println("Insira os valores de x e y")

fmt.Scan(&x, &y)

if y < 0 || x < 0{
fmt.Println("Erro, raiz de numero negativo")
} else if math.Pow(y, 3) - math.Sqrt(x) == 0 {
fmt.Println("Erro, divisao por zero")
} else  {
s := (math.Pow(x, 3)-math.Sqrt(y)) / (math.Pow(y,3)-math.Sqrt(x))
fmt.Printf("f(x,y) = %.4f", s)
}
}
