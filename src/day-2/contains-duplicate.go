package day2

import "fmt"

/*
🧩 Desafio 3: Contains Duplicate
Problema: Dado um array de inteiros, retorne true se algum valor aparecer mais de uma vez, e false se todos forem distintos.
Input:  [1, 2, 3, 1]
Output: true

Input:  [1, 2, 3, 4]
Output: false

Input:  [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
Output: true
Constraints:

1 <= len(nums) <= 10⁵
-10⁹ <= nums[i] <= 10⁹

*/

func RunContainsDuplicate() {
	fmt.Println("Contém número duplicado?", containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	prevNums := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		_, exists := prevNums[num]

		if exists {
			return exists
		}

		prevNums[num] = struct{}{}
	}

	return false
}
