package main

import "fmt"

func main() {

var v, sum,avg float64
var p float64

fmt.Println("Insira 6 valores")

for i := 0; i < 6; i++ {

fmt.Scan(&v)

if v > 0 {
p += 1
sum += v
}


}

fmt.Println(p)

avg = sum/p

fmt.Println(avg)

}
