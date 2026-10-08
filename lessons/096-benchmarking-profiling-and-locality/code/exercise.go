package bench

func Sum(values []int) int {
	// TODO: implement the baseline and add a benchmark comparing a locality-friendly variant.
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
