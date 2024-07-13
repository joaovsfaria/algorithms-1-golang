package main

import "fmt"

func main() {

var v0, s0, a, t float64

fmt.Println("Insira os valores da v0, s0, a, t")

fmt.Scan(&v0, &s0, &a, &t)

v := v0 + a * t
s := s0 + v0*t + 0.5 * a * t * t
fmt.Printf("%v %v", v, s)

}
