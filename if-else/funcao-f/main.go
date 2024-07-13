package main

import "fmt"
import "math"

func main() {

var x, y, f float64

fmt.Println("Insira os valores de x e y")
fmt.Scan(&x, &y)

if x > y {
f = math.Pow(x, 2) - math.Pow(y, 2) + 2 * x * y
fmt.Printf("f(x,y) %v", f)
} else if x == y {
f = 2 * x * y + x + y
fmt.Printf("f(x,y) = %v", f)
} else {
f = math.Pow(y, 2) + math.Pow(x, 2) + 2 * x * y
fmt.Printf("f(x,y) = %v", f)
}
}
