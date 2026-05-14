package dayOne

import "fmt"

func DayOne() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println("Indíces da soma:", twoSum(nums, target))

	fmt.Println("O 'parenteses' é válido?", validateParentheses("()[]{}"))

	fmt.Println("O melhor lucro é:",
		bestTimeToBuy([]int{7, 1, 5, 3, 6, 4}),
	)
}
