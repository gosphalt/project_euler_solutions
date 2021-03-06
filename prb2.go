/*

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

https://projecteuler.net/problem=2

*/

package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func checkIfMod(num uint64, numtype string) uint64{
  var testmod uint64
  if numtype == "EVEN"{
    testmod = 0
  }else{
    testmod = 1
  }

  if num % 2 == testmod {
    return num
  }else{
    return 0
  }
}

func main(){
    args := os.Args
    if (len(args) != 3 ) || (len(args) == 3 && strings.ToUpper(args[2]) != "ODD" && strings.ToUpper(args[2]) != "EVEN" && strings.ToUpper(args[2]) != "ALL" ) {
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [limit] [even/odd/all]")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 1000 even")
      return
    }

    var sum uint64
    var fib_cur, fib_prev, fib_next uint64
    var limit uint64
    var typenum string

    limit,_ = strconv.ParseUint(args[1],10,64)
    typenum = strings.ToUpper(args[2])

    fib_cur = 2
    fib_prev = 1

    sum += checkIfMod(fib_prev, typenum)
    sum += checkIfMod(fib_cur, typenum)

    for (fib_prev + fib_cur) < limit{
      fib_next = fib_prev + fib_cur
      sum += checkIfMod(fib_next, typenum)
      fib_prev = fib_cur
      fib_cur = fib_next
    }

    fmt.Printf("The sum is %v\n",sum)
}
