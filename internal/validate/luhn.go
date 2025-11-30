package validate

func IsValidLuhn(cardNum string) bool {
	if len(cardNum) > 19 || len(cardNum) < 13 {
		return false
	}
	n := len(cardNum)
	sum := 0
	double := false

	for i := n - 1; i >= 0; i-- {
		c := cardNum[i]
		d := int(c - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}
