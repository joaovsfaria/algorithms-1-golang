package main

import "fmt"

func main() {

var pv1, pv2, ep float64

fmt.Println("Insira as notas da PV1 e PV2e a media dos exercicios propostos")

fmt.Scan(&pv1, &pv2, &ep)

mf := (((pv1+pv2)/2)*0.9) + ep*0.1

fmt.Printf("%.2f", mf) 

}
