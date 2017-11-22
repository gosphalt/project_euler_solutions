/*

The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

https://projecteuler.net/problem=6

*/

package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
  "math"
)

func main(){
    start := time.Now()

    args := os.Args
    if (len(args) != 3 ) {
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [number] [number]")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 2 100 ")
      return
    }

    power,_ := strconv.Atoi(args[1])
    numbers,_ := strconv.Atoi(args[2])

    sum_power := uint64(0)
    sum := uint64(0)

    for i:=1;i<=numbers;i++{
      sum_power += uint64(math.Pow(float64(i),float64(power)))
      sum += uint64(i)
    }


    diff := uint64(math.Abs(float64(sum_power) - math.Pow(float64(sum),float64(power))))

    fmt.Printf("Duration: %v\n",time.Now().Sub(start))
    fmt.Printf("The difference between the sum of the %v-power of the first %v natural numbers and the %v-power of the sum is %d\n",power, numbers, power, diff)
}
