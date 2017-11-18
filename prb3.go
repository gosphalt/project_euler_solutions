/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

https://projecteuler.net/problem=3

*/

package main

import (
  "fmt"
  "os"
  "math"
  "strconv"
)

type PrimeArray []uint64

func testPrimeNumber(num uint64, primes PrimeArray, primes_ptr *PrimeArray) bool{
  for i:=range primes{
    if num%primes[i]==0{
      return false
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
      return false
    }
  }

  *primes_ptr = append(primes,num)
  return true
}

func main(){
    args := os.Args
    if (len(args) != 2 ) {
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [number]")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 600851475143")
      return
    }

    var primes PrimeArray
    num,_ := strconv.ParseUint(args[1],10,64)
    lpf := uint64(0)

    for i:=uint64(2);i <= uint64(num/2);i++{
      if num%i==0 && testPrimeNumber(i,primes,&primes) && i >= lpf{
        lpf = i
      }
    }

    fmt.Printf("The largest prime factor is %v\n",lpf)
}
