package main

import (
	"fmt"
	"github.com/schwiet/project-euler-go/utils"
)

/*
   A palindromic number reads the same both ways. The largest palindrome made
   from the product of two 2-digit numbers is 9009 = 91 × 99.

   Find the largest palindrome made from the product of two 3-digit numbers.
*/

func main() {
	result := lpp(100000, 999999)

	// should print 906609
	fmt.Println("Largest Palindrome Product:", result)
}

func lpp(min, max int) int {
	var product int
	max_p := 0

	for num_a := max; num_a > min; num_a -= 1 {
		// my thoughts are that this is a reliable way to short-circuit the
		// result and prevent looking much farther than we need to. If we get to
		// a point where none of the remaining products are greater than the
		// product we've found, than continuing to look is pointless.
		// TODO: maybe there is an earlier point to know we're done...
		if max*num_a <= max_p {
			return max_p
		}

		// each iteration, num_b only iterates down to the current num_a value
		// this avoids performing duplicate calculations of each product
		for num_b := max; num_b >= num_a; num_b -= 1 {
			product = num_a * num_b
			if product == utils.ReverseDigits(product) && product > max_p {
				max_p = product
			}
		}
	}

	return max_p
}
