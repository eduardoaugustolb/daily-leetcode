package day1

import "fmt"

/*
 Two Sum
 Dado um array de inteiros nums e um inteiro target, retorne os índices dos dois números que somam ao target.
 Cada entrada tem exatamente uma solução. Não use o mesmo elemento duas vezes.

 Exemplos
 Input:
 nums = [2,7,11,15]
 target = 9

 Output: [0,1]
 Input:
 nums = [3,2,4]
 target = 6

 Output: [1,2]
*/

func RunTwoSum() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println("Indíces da soma:", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// num: idx
	prevNumsIdx := make(map[int]int, len(nums))
	for idx, num := range nums {
		diff := target - num
		anotherNumberIdx, anotherNumberIdxExists := prevNumsIdx[diff]
		if anotherNumberIdxExists {
			return []int{anotherNumberIdx, idx}
		}
		prevNumsIdx[num] = idx
	}

	return []int{}
}
