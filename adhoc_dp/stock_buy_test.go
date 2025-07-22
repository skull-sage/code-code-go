package adhoc_dp

import (
	"testing"
)

func maxProfit(prices []int) int {
	// we simply care about the profit against current option
	low := prices[0]
	profit := 0

	for idx := 1; idx < len(prices); idx++ {
		curr := prices[idx]
		if curr > low {
			if curr-low > profit {
				profit = curr - low
			}
			//fmt.Println("low", low, "curr", curr, "=>", profit)
		} else {
			low = curr
		}
	}

	return profit

}

func TestStockBuy(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
}
