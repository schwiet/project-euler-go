package main

import (
    "fmt"
    "github.com/schwiet/project-euler-go/utils"
)

/*
    The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

    Find the sum of all the primes below two million.
 */

const Limit = 2000000;

func main() {

    // start with a value of two, to account for 2 being a prime number, but not
    // acounted for in the loop below
    sum := 2

    // this loops thru all odd numbers greater than 2
    for i := 3; i < Limit; i += 2 {
        if utils.IsPrime( i ){
            sum += i
        }
    }

    // should print 142913828922
    fmt.Println( "sum of first 2m primes:", sum )
}
