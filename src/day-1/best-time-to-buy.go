package day1

/*
Best Time to Buy and Sell Stock
Dado um array prices onde prices[i] é o preço de uma ação no dia i, retorne o lucro máximo que você pode obter comprando em um dia e vendendo em um dia futuro. Se não for possível lucrar, retorne 0.


Você só pode comprar uma vez e vender uma vez. Você deve comprar antes de vender.

Exemplos
Input:  [7,1,5,3,6,4]
Output: 5
// compra dia 1 (preço 1)
// vende dia 4 (preço 6)
Input:  [7,6,4,3,1]
Output: 0
// preços só caem
// nenhuma transação
*/

func bestTimeToBuy(prices []int) int {
	lowestPrice := prices[0]
	bestProfit := 0
	for _, price := range prices {
		if price < lowestPrice {
			lowestPrice = price
		}
		profit := price - lowestPrice
		if bestProfit < profit {
			bestProfit = profit
		}
	}
	return bestProfit
}
