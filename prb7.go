/*

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

https://projecteuler.net/problem=7

*/

package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
  "math"
)

type PrimeArray []uint64

func testPrimeNumber(num uint64, primes PrimeArray) PrimeArray{
  for i:=range primes{
    if num%primes[i]==0{
      return primes
    }
  }

  var start uint64
  limit := uint64(math.Sqrt(float64(num)))
  if len(primes) > 0{
    start = primes[len(primes)-1]+1
  }else{
    start = uint64(2)
  }

  for i:=start;i<limit;i++{
    if num%i==0{
      return primes
    }
  }

  primes_out := append(primes, num)
  return primes_out
}

func main(){
    start := time.Now()

    args := os.Args
    if (len(args) != 2 ) {
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [number] ")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 1001 ")
      return
    }

    var primes PrimeArray
    num_primes,_ := strconv.ParseUint(args[1],10,64)

    i := uint64(2)

    for ;uint64(len(primes)) < uint64(num_primes);{
        primes = testPrimeNumber(i, primes)
        i++
    }

    fmt.Printf("Duration: %v\n",time.Now().Sub(start))
    fmt.Printf("The %vth prime number is %d\n",num_primes, primes[len(primes)-1])
}
