package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(maximumSubarraySum([]int{0, 5, 0, 10, 0}, 4))                     // Output: 3")
	fmt.Println(maximumSubarraySumWithNegativeNumbers([]int{-10, -5, -2, -8}, 2)) // Output: 9
}

/*
1. Soma Máxima de um Subarray de Tamanho Fixo (Easy)
Problema: Dado um array de inteiros e um número inteiro k, encontre a soma máxima de um subarray contíguo de tamanho k.

Exemplo:

Array: [2, 1, 5, 1, 3, 2]

k=3

Subarrays de tamanho 3 e suas somas:

[2, 1, 5] -> Soma: 8

[1, 5, 1] -> Soma: 7

[5, 1, 3] -> Soma: 9

[1, 3, 2] -> Soma: 6

Soma máxima: 9

*/
//Array: [2, 1, 5, 1, 3, 2]
func maximumSubarraySum(arr []int, k int) int {
	if len(arr) < k {
		return 0 // Retorna 0 se o array for menor que k
	}
	currentSum := 0
	maxSum := 0
	windowEndIndex := 0

	for windowEndIndex < len(arr) {
		if windowEndIndex < k {
			currentSum += arr[windowEndIndex]
		} else {
			// Quando a janela atinge o tamanho k, soma o atual elemento da janela e subtrai o primeiro elemento da janela
			currentSum += arr[windowEndIndex] - arr[windowEndIndex-k]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}

		windowEndIndex++
	}

	return maxSum
}

func maximumSubarraySumWithNegativeNumbers(arr []int, k int) int {
	if len(arr) < k || k <= 0 { // Adicionando a validação para k <= 0
		return 0
	}

	currentSum := 0
	maxSum := math.MinInt32 // Ou math.MinInt64 se os valores puderem ser muito grandes/pequenos
	windowEndIndex := 0

	for windowEndIndex < len(arr) {
		if windowEndIndex < k-1 {
			currentSum += arr[windowEndIndex]
		} else {
			currentSum += arr[windowEndIndex]
			if windowEndIndex >= k {
				currentSum -= arr[windowEndIndex-k]
			}
		}
		if windowEndIndex >= k-1 {
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
		windowEndIndex++
	}

	return maxSum
}
