package main

func shippingCost(weight int) int {
	if weight <= 0 {
		return 0
	}
	return 500 + weight*100
}
