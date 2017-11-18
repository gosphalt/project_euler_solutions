/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

https://projecteuler.net/problem=1

*/

package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func testMod(num int, divs[]int) bool{
  for i := range divs{
    if num%divs[i]==0{
        return true
    }
  }
  return false
}

func main(){
    args := os.Args
    if len(args) != 3{
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [limit] [divisors]")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 1000 3,5")
      return
    }

    sum := 0
    limit,_ := strconv.Atoi(args[1])
    divisors_str := strings.Split(args[2],",")
    var divisors[]int

    for i:= range divisors_str{
      intval,_ := strconv.Atoi(divisors_str[i])
      divisors = append(divisors, intval)
    }

    for i:=1;i<limit;i++{
      if testMod(i,divisors){
        sum+= i
      }
    }

    fmt.Printf("The sum is %v\n",sum)
}
