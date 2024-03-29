package main

import (
	"fmt"
)

/*
  Starting in the top left corner of a 2×2 grid, and only being able to move to
  the right and down, there are exactly 6 routes to the bottom right corner.

  How many such routes are there through a 20×20 grid?
*/

const GRID_SIZE uint64 = 20

func main() {

	// paths := brute(GRID_SIZE)
	// paths := calcPerms(GRID_SIZE)
	paths := createAllPerms(GRID_SIZE)
	fmt.Println("For a Grid of", GRID_SIZE, ", there are", paths, "paths.")
}

func createAllPerms(gridSize uint64) uint64 {
	downs := gridSize
	rights := gridSize

	cache := pathMemo{make(map[uint64]map[uint64]uint64)}
	permutations := cache.pathsFrom(downs, rights)
	return permutations
}

type pathMemo struct {
	cached map[uint64]map[uint64]uint64
}

func (p *pathMemo) pathsFrom(remDown, remRight uint64) uint64 {
	// first, check for cached value
	r, ok := p.cached[remDown]
	if ok {
		c, ok := r[remRight]
		if ok {
			return c
		}
	} else {
		r = make(map[uint64]uint64)
		p.cached[remDown] = r
	}
	//
	var sum uint64 = 0

	// reached the end
	if remDown == 0 && remRight == 0 {
		return 1
	}

	// if there can be any more steps down, make recursive call
	if remDown > 0 {
		sum += p.pathsFrom(remDown-1, remRight)
	}

	// if there can be any more steps right, make recursive call
	if remRight > 0 {
		sum += p.pathsFrom(remDown, remRight-1)
	}

	// cache the sum calculated for this grid: remDown, remRight
	r[remRight] = sum
	return sum
}

/*
 * CAN'T BE CALCULATED WITH 64 bits
 */

func calcPerms(gridSize uint64) uint64 {
	steps := gridSize * 2
	n := factorialWithRange(steps, gridSize)
	r := factorial(gridSize)
	fmt.Println(n, r)
	return n / r
}

func factorialWithRange(start, end uint64) uint64 {

	var value uint64 = 1
	var i uint64
	for i = start; i > end; i -= 1 {
		value *= i
	}
	return value
}

func factorial(n uint64) uint64 {
	return factorialWithRange(n, 0)
}

/*
 * TAKES TOO LONG
 */

func brute(gridSize uint64) uint64 {
	var steps, paths uint64
	// the number of steps is the sum of the sides
	steps = gridSize * 2

	var i, j, rtCt, dnCt uint64
	sequence := make([]int, steps)

	// this is just a brute force, binary sequence generator. There must be a
	// more efficient way, but this is the solution that first occured to me.
	for i < steps {
		if sequence[0] == 0 {
			sequence[0] = 1
		} else {
			for i = 0; i < steps && sequence[i] != 0; i += 1 {
				sequence[i] = 0
			}
			if i < steps {
				sequence[i] = 1
			}
		}

		// loop thru the sequence and count the 1's and 0's. if the counts are
		// equal, it represents a valid path (it moves DOWN and RIGHT the same
		// number of times, so it ends up in the opposite corner of the grid)
		dnCt = 0
		rtCt = 0
		for j = 0; dnCt*2 <= steps && rtCt*2 <= steps && j < steps; j += 1 {
			if sequence[j] == 1 {
				dnCt += 1
			} else {
				rtCt += 1
			}
		}
		if j == steps && dnCt == rtCt {
			paths += 1
			// fmt.Println(sequence, oneCt, zroCt, steps)
		}
	}

	return paths
}
