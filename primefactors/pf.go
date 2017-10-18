package primefactors

func GeneratePrimeFactors(value int) []int {
	rval := []int{}
	for divisor := 2; value > 1; divisor++ {
		for ; value%divisor == 0; value /= divisor {
			rval = append(rval, divisor)
		}
	}
	return rval
}
