package main

import "fmt"
import "math"


func main() {

var r, r1, r2, r3, v float64


fmt.Println("Insira os valores da tensao de r1, r2, r3")

fmt.Scan(&v, &r1, &r2, &r3)

i := v*((1/r1)+(1/r2)+(1/r3))

r = 1/(math.Pow(r1, -1) + math.Pow(r2, -1) + math.Pow(r3,-1))



fmt.Printf("%.2f %.2f", i, r)

}
