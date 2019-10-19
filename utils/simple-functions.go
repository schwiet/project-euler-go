package utils

import (
    "math"
)

/*
    IsPrime returns true if the passed num is prime, otherwise it returns false.
    It is implemented with a simple primality algorithm
*/
func IsPrime(num int) bool{

    if num <= 3 {
        return num > 1
    } else if num % 2 == 0 || num % 3 == 0 {
        return false
    } else{
        // every prime greater than 6 is of the form 6k + 1 or 6k - 1
        for i := 5; i * i <= num; i += 6 {
            if num % i == 0 || ( num % ( i + 2 ) ) == 0 {
                return false
            }
        }
        return true
    }
}

/*
    GetFactors returns all of the factors of num
*/
func GetFactors(num int) []int{
    var result []int
    limit := int( math.Floor( math.Sqrt( float64( num ) ) ) )

    // loop from 1 to the square root on num
    for i := 1; i < limit; i += 1{
        // if num is divisible by i, add i and num/i to the results
        if num % i == 0 {
            result = append( result, i )
            // don't add num/i if it is equal to i (i.e i == sqrt(num))
            if i != limit {
                result = append( result, num/i )
            }
        }
    }

    return result
}

/*
    ReverseDigits takes num and returns an integer whose digits are in the
    opposite order: e.g 12345 -> 54321
 */
func ReverseDigits(num int) int {
    var digit int
    var result, next, current int = 0, num, num

    for ; next > 0; {
        // get the next digit by subtracting the current number from itself, but
        // with the loss of percision of the ones place
        next /= 10
        digit = ( current - ( next * 10 ) )
        // ready for next iteration
        current = next

        // multiply the result by 10 and add the single digit
        result = result * 10 + digit
    }

    return result
}

/*
    Multiply each item in the slice to get the product
 */
func GetSliceProduct( slice []int ) int{
    product := 1
    for _, v := range slice {
        product *= v
    }
    return product
}
