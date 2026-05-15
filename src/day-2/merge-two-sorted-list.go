/*
🧩 Desafio 1: Merge Two Sorted Lists
Problema: Dados os heads de duas listas ligadas ordenadas, merge elas em uma única lista ordenada. Retorne o head da lista merged.
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = [0]
Output: [0]
Constraints: 0 <= N <= 50, valores entre -100 e 100
*/

package day2

func mergeTwoSortedList(list1 []int, list2 []int) []int {
	i := 0
	j := 0
	k := 0
	subArrSize := len(list1) + len(list2)
	subArr := make([]int, subArrSize)
	for i < len(list1) && j < len(list2) && k < subArrSize {
		list1Item := list1[i]
		list2Item := list2[j]

		if list1Item < list2Item {
			subArr[k] = list1Item
			i++
		} else {
			subArr[k] = list2Item
			j++
		}

		k++

	}

	list1Rest := list1[i:]
	copy(subArr[k:], list1Rest)
	k += len(list1Rest)
	copy(subArr[k:], list2[j:])

	return subArr
}
