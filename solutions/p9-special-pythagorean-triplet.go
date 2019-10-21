package main

import (
    "math"
    "fmt"
)

/*
    A Pythagorean triplet is a set of three natural numbers, a < b < c, for
    which,
        a^2 + b^2 = c^2

    For example, 3^2 + 4^2 = 9 + 16 = 25 = 52.

    There exists exactly one Pythagorean triplet for which a + b + c = 1000.
    Find the product abc.
 */


func main() {
    result := findPythagoreanTripletProduct( 3.0, 4.0, 1000.0 )

    // should print 31875000
    fmt.Println( "Pythagorean Triplet Product with Sum:", 1000, "is:", result )
}

func findPythagoreanTripletProduct( a, b, sum float64 ) float64 {
    for {
        // find pythagorean c of a & b
        c := math.Sqrt( math.Pow( a, 2 ) + math.Pow( b, 2 ) )

        // success case
        if a + b + c == sum {
            return ( a * b * c )

        // if a has reached a number that ensures we have exceeded the sum we
        // need, there is no such triplet
        } else if a > sum / 3{
            return -1.0

        // our triplet has exceeded the sum, so notch up a and reset b to a+1
        } else if a + b + c > sum {
            a += 1
            b = a + 1

        // increment b
        } else {
            b += 1
        }
    }
}
