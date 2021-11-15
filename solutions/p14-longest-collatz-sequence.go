package main

import "fmt"

/*
  The following iterative sequence is defined for the set of positive integers:

  n → n/2 (n is even)
  n → 3n + 1 (n is odd)

  Using the rule above and starting with 13, we generate the following sequence:

  13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
  It can be seen that this sequence (starting at 13 and finishing at 1) contains
  10 terms. Although it has not been proved yet (Collatz Problem), it is thought
  that all starting numbers finish at 1.

  Which starting number, under one million, produces the longest chain?

  NOTE: Once the chain starts the terms are allowed to go above one million.
*/

const RANGE uint64 = 1000000

func main() {
	var i, n, seq_len, max_len, starting uint64
	for i = 1; i <= RANGE; i += 1 {
		// reset the sequence length counter and n to next starting number
		seq_len = 0
		n = i

		// apply the Collatz sequence rules until n == 1
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			seq_len += 1
		}

		// determine if the sequence length is larger than any we've seen
		// if so, keep a reference
		if seq_len > max_len {
			max_len = seq_len
			starting = i
		}
	}

	fmt.Println("Answer:", starting, "Sequence Length:", seq_len)
	// Answer: 837799 Sequence Length: 152
}
