package day2

import "fmt"

func DayTwo() {
	fmt.Println("Fazer merge e sort de duas listas")
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}
	fmt.Println("Antes:", list1, list2)
	fmt.Println("Depois", mergeTwoSortedList(list1, list2))

	numsList := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maior soma do array:", numsList, maxSubArray(numsList))

	fmt.Println("Contém número duplicado?", containsDuplicate([]int{1, 2, 3, 1}))
}
