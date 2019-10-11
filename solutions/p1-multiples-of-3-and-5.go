package main

import (
    "fmt"
)

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
