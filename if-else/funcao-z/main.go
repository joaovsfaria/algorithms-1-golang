package main 

import "fmt"
import "math"


func main() {

var x, y float64

fmt.Println("Insira x, y")
fmt.Scan(&x, &y)

if x + y <= 0 || x - y < 0{
fmt.Print("Calculo impossivel")
} else {
z := math.Sqrt(x-y)/math.Sqrt(x+y)
fmt.Println(z)
}
}
