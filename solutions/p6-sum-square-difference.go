package main

import (
    "fmt"
    "math"
)
/*
    The sum of the squares of the first ten natural numbers is,
        1^2 + 2^2 + ... + 10^2 = 385

    The square of the sum of the first ten natural numbers is,
        (1 + 2 + ... + 10)^2 = 55^2 = 3025

    Hence the difference between the sum of the squares of the first ten natural
    numbers and the square of the sum is 3025 − 385 = 2640.

    Find the difference between the sum of the squares of the first one hundred
    natural numbers and the square of the sum.
 */

func main(){
    var sum_of_sq, sum float64 = 0, 0

    for i := 1.0; i <= 100; i += 1 {
        sum += i
        sum_of_sq += math.Pow( i, 2 )
    }

    sq_of_sums := math.Pow( sum, 2 )

    // should print: 2.516415e+07
    fmt.Println( "Sum Square Dif:", sq_of_sums - sum_of_sq )
}
