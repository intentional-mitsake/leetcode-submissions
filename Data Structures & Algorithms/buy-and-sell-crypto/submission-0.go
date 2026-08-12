func maxProfit(prices []int) int {
    maxProfit := 0
	minBuy := 99999999999999 // could just use math.MaxInt32
	for _, price := range prices {
		// if the price is less, use that
		if price < minBuy {
			minBuy = price
		}
		// if profit better, save it
		if price - minBuy > maxProfit {
			maxProfit = price - minBuy
		} 
	}
	return maxProfit
}
