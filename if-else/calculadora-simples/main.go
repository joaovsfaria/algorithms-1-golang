package main

import "fmt"

func main() {

var v1, v2, res float64
var op rune


fmt.Println("Insira dois valores")
fmt.Scan(&v1, &v2)

fmt.Println("Insira o operador: 1 - Adição; 2 - Subtração; 3 - Multiplicação; 4 - divisão")
fmt.Scan(&op)

switch op {
case 1:
res = v1 + v2
fmt.Println(res)
case 2:
res = v1 - v2
fmt.Println(res)
case 3:
res = v1 * v2
fmt.Println(res)
case 4:
if v2 == 0 {
fmt.Println("Erro")
} else {
res = v1 / v2
}
default:
fmt.Println("Opção não existe")
}



}
