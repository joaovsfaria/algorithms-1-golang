package main

import "fmt" 
import "math"

func main() {

var t, s float64

fmt.Println("Insira o valor de t")
fmt.Scan(&t)

if t >= -1 && t <= 1 { 
s = (1 - math.Abs(t))
} else {
s = 0
}

fmt.Printf("x(t) = %v", s)

}






