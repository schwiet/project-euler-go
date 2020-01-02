package main

import (
    "fmt"
    "github.com/schwiet/project-euler-go/utils"
)

const minimum = 501

func main(){
    factors := 0
    tnum := 0
    for i := 1; factors < minimum; i += 1 {
        tnum += i
        factors = len( utils.GetFactors( tnum ) )
    }

    //Answer: 76576500
    fmt.Println( "The First Triangle Number with at least", minimum, " factors, is:", tnum )
}
