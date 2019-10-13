package main

import (
    "fmt"
    "github.com/schwiet/project-euler-go/utils"
)

/*
    The prime factors of 13195 are 5, 7, 13 and 29.

    What is the largest prime factor of the number 600851475143 ?
*/

func main(){

    // get all of the factors of the number
    factors := utils.GetFactors( 600851475143  )
    // to hold the value of the highest prime factor
    max_prime_f := 1

    // loop thru the factors, if the value is prime and it is larger than the
    // highest prime factor we have seen, so far, store it.
    for _, v := range factors {
        if utils.IsPrime( v ) && v > max_prime_f {
            max_prime_f = v
        }
    }

    // should print 6857
    fmt.Println( "max prime factor?", max_prime_f )
}
