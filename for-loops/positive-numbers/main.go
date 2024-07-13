package main

import "fmt"

func main() {

var v1, quantidade int

fmt.Println("insira 6 valores")


for i := 0; i < 6; i++ {
fmt.Scan(&v1)

if v1 > 0 {
quantidade++
}
}

fmt.Println(quantidade)



}
