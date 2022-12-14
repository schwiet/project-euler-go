package main

import (
	"fmt"
)

/*
	215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
	What is the sum of the digits of the number 21000?
*/

func main() {
	// TODO: use os.Args as input
	// if len(os.Args) != 3 {
	// 	fmt.Println("Must provide two args: base, power")
	// 	return
	// }

	result := powerDigitSum(2, 1000)
	fmt.Println("Result:", result)
}

func powerDigitSum(base, power int) int {
	var digits []int
	overflow := make([]int, power)
	var result int
	var firstOverflowPos int
	var product int
	var ovf int
	// the first digit equals base
	dig := base
	// first pass needs to imply overflow (so first loop initiates)
	hasOverflowed := true

	// successively multiply each product by <base>. keep track of any time the
	// product is greater than or equal to 10 (overflows beyond a single digit),
	// noting the multiplication iteration. If there were any such occurrences,
	// repeat the multiplication loop adding overflowing digits from the previous
	// pass. Do this until a pass produces no overflows. This allows us to
	// calculate a huge string of digits, avoiding having to store the product in
	// a bit-bound type. The tradeoff is efficiency; this multiplication loop has
	// to start over any time the previous iteration had products that overflowed
	for i := 0; hasOverflowed; i += 1 {
		hasOverflowed = false
		// power - 1 because x ^ 1 = x
		for j := firstOverflowPos; j < power-1; j += 1 {
			// mulitply the current digit by the base number, then add any overflow
			// from the previous place
			product = dig*base + overflow[j]
			// determine the next digit
			dig = product % 10
			// if the product overflows into the next place, keep the value for the
			// next outer iteration to add to it's product at the corresponding step
			ovf = product / 10
			overflow[j] = ovf
			// this is just an optimization so we can skip any multiply-by-zeros in
			// the next iteration
			if ovf > 0 && !hasOverflowed {
				firstOverflowPos = j
				hasOverflowed = true
			}
		}
		// have multiplied by base <power> times for this place, so we have the
		// final digit for this place. Add it to the result list and add the digit
		// to the running tally
		digits = append(digits, dig)
		result += dig
		// reset
		dig = 0
	}

	// TODO: print the digits nicely
	fmt.Println("Digits", digits)
	return result
}
