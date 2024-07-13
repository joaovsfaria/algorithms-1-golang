package main

import (
	"fmt"
	"math"
)

func main() {

	var x, m, s float64

	fmt.Println("Insira o valor de x, m, s")

	fmt.Scan(&x, &m, &s)

	f := (1 / math.Sqrt(2*math.Pi*math.Pow(s, 2)) * math.Pow(math.E, -(math.Pow((x-m), 2)/(2*math.Pow(s, 2)))))

	fmt.Println(f)

}
