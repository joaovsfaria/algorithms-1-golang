//Escreva um programa que leia o numero de um funcionario, seu numero de horas trabalhadas, o valor que recebe por 
// hora e calcula o salario desse funcionario. A seguir, mostre o numero e o salario do funcionario.


package  main

import "fmt"

func main() {

var numeroFuncionario int
var numeroHorasTrabalhadas float64
var valorHorasTrabalhadas float64


fmt.Println("Insira o numero, a quantidade de horas trabalhadas e o valor recebido por hora pelo funcionario?")
fmt.Scan(&numeroFuncionario)
fmt.Scan(&numeroHorasTrabalhadas)
fmt.Scan(&valorHorasTrabalhadas)

salario := numeroHorasTrabalhadas * valorHorasTrabalhadas

fmt.Printf("NUMBER = %v\n", numeroFuncionario)
fmt.Printf("SALARY = U$ %.2f", salario)





}
