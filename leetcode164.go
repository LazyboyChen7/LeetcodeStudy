func fractionToDecimal(numerator int, denominator int) string {

	symbol := ""

	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0){

		symbol = "-"

	}

	numerator = abs(numerator)

	denominator = abs(denominator)

	integer := numerator/denominator

	numerator = numerator%denominator

	if numerator == 0 {

		return fmt.Sprintf("%s%d", symbol, integer)

	}

	num := []int{}

	m := map[int]int{}

	idx := -1

	for numerator > 0 {

		if _, ok := m[numerator]; ok {

			idx = m[numerator]

			break

		}

		m[numerator] = len(num)

		numerator *= 10

		inter := numerator/denominator

		numerator = numerator%denominator

		num = append(num, inter)

	}

	d := ""

	for i := 0; i < len(num); i++ {

		if i != idx {

			d += fmt.Sprintf("%d", num[i])

		} else {

			d += fmt.Sprintf("(%d", num[i])

		}

	}

	if idx != -1 {

		d += ")"

	}

	return fmt.Sprintf("%s%d.%s", symbol, integer, d)

}

func abs(x int) int {

	if x < 0 {

		x *= -1

	}

	return x

}