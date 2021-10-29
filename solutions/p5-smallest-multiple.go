package main

/*
   2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

   What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

import (
	"fmt"
	"math"
)

func main() {
	result := sm(1, 20)

	// should print 232792560
	fmt.Println("Smallest multiple of 1 thru 20", result)
}

func sm(min, max uint) uint {

	// TODO: could determine all of the numbers that are factors of max and
	// also skip those in the inner loop below.

	// take the max number an increment each candidate by that (since every
	// candidate must be a multiple of min-max)
	for i := max; i < math.MaxUint32; i += max {
		// initialize flag
		failed := false
		// try dividing our candidate, i, by each number in the range. We can
		// skip max since we already know it's a multiple of max (we're
		// incrementing by max in the outer loop)
	DivByEach:
		for j := min; j < max; j += 1 {
			if i%j != 0 {
				failed = true
				break DivByEach
			}
		}

		// if each number divided evenly, we found the smallest multiple
		if !failed {
			return i
		}
	}

	return 0
}
