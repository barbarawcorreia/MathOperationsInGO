package main

import "fmt"

// ARRAY MATEMÁTICO
// array a, b com números inteiros
// array c com sinais
// array d com resultado

func main() {

}

func MathOperations(numbersA []int, numbersB []int, operations []string) []int {
	result := []int{}

	for index, eachOperation := range operations {
		fmt.Println("OPERATION", eachOperation)
		fmt.Println("numbersA[index]:", numbersA[index])
		fmt.Println("numbersB[index]", numbersB[index])
		fmt.Println("INDEX", index)
		fmt.Println("////////////////////////////////")

		if eachOperation == "+" {
			result = append(result, numbersA[index]+numbersB[index])
		} else if eachOperation == "-" {
			result = append(result, numbersA[index]-numbersB[index])
		} else if eachOperation == "*" {
			result = append(result, numbersA[index]*numbersB[index])
		}
	}

	return result
}
