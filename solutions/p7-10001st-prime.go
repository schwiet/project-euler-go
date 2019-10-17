package main

import (
    "fmt"
    "github.com/schwiet/project-euler-go/utils"
)

/*
    By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
    that the 6th prime is 13.

    What is the 10 001st prime number?
 */

func main(){

    var i int = 3;

    // the only obvious optimization was to not check even numbers (above 2),
    // so we start at 3 and initialize count to 1, since 2 is prime.
    for count := 1; count < 10001; i += 2 {
        if utils.IsPrime( i ){
            count += 1
        }
    }

    // should print: 104743
    fmt.Println( "10001st Prime:", i - 2 );
}
