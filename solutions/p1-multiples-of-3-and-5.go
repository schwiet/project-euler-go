package main

import (
    "fmt"
)

/*
    If we list all the natural numbers below 10 that are multiples of 3 or 5, we
    get 3, 5, 6 and 9. The sum of these multiples is 23.

    Find the sum of all the multiples of 3 or 5 below 1000.
 */
 
func main(){
    sum := 0
    // sum all of the multiples of 3
    for i := 3; i < 1000; i += 3 {
        sum += i
    }
    // sum all of the multiples of 5, as long as they are not common multiples
    // of 3
    for i := 5; i < 1000; i += 5 {
        // don't count the same number twice
        if i % 3 != 0 {
            sum += i
        }
    }
    // should print	233168
    fmt.Println( sum )
}
