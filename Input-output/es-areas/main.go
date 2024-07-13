package main

import "fmt"

func main() {

var totalDinheiro float64
var precoTotal float64

fmt.Println("Insira o valor total de dinheiro levado por Zezinho e o preco total da compra")

fmt.Scan(&totalDinheiro)
fmt.Scan(&precoTotal)

troco := totalDinheiro - precoTotal

fmt.Printf("O troco é de R$%.2f", troco)

}
