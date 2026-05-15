package day2

/*
🧩 Desafio 2: Maximum Subarray
Problema: Dado um array de inteiros, encontre o subarray contíguo (sequência de elementos adjacentes) que tem a maior soma e retorne essa soma.
Input:  [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6
Motivo: [4, -1, 2, 1] são elementos adjacentes cuja soma é máxima
Input:  [1]
Output: 1
Input:  [-2, -1]
Output: -1  ← quando tudo é negativo, retorna o menos negativo
Constraints:

Array de inteiros (pode ter negativos)
1 <= len(nums) <= 10⁵
Retorna um int
*/

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	bestSum := currentSum

	for _, num := range nums {
		currentSum = max(num, currentSum+num)

		if currentSum > bestSum {
			bestSum = currentSum
		}
	}

	return bestSum

}
