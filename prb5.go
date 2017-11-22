/*

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

https://projecteuler.net/problem=4

*/

package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
)

func main(){
    start := time.Now()

    args := os.Args
    if (len(args) != 3 ) {
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [number] [number]")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 1 20 ")
      return
    }

    first_factor,_ := strconv.Atoi(args[1])
    last_factor,_ := strconv.Atoi(args[2])

    spn_check := uint64(last_factor)-1
    spn_found := false

    for ;!spn_found;{
        spn_check = spn_check + 1
        i := uint64(first_factor)
        for ; i<=uint64(last_factor); {
            if spn_check%i != 0 {
              spn_found  = false
              break
            }else{
              spn_found = true
            }
            i = i+1
        }
    }

    fmt.Printf("Duration: %v\n",time.Now().Sub(start))
    fmt.Printf("The smallest positive number that is evenly divisible by all of the numbers from %v to %v is %v\n",first_factor, last_factor, spn_check)
}
