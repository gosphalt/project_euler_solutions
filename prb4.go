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
  "math"
  "time"
)

func testPalindrome(num uint64) bool{
  num_str := strconv.FormatUint(num,10)

  for i:=0;i< len(num_str)/2;i++{

    if num_str[0+i:0+i+1] != num_str[len(num_str)-i-1:len(num_str)-i]{
      return false
    }
  }
  //fmt.Printf("%v is PALYNDROME!\n",num)
  return true
}

func checkBreak(arr []uint64, numcheck uint64) bool{
  check := true
  for i := range arr{
    if arr[i] != numcheck {
      check = false
    }
  }
  return check
}

func modArray(arrInput []uint64, index uint64, limitNum uint64) []uint64{
  arrOutput := arrInput

  if checkBreak(arrInput, limitNum-1){
    return arrOutput
  }

  arrOutput[index] = arrOutput[index] + 1

  if arrOutput[index]%limitNum == 0{
      arrOutput[index] = 1
      arrOutput = modArray(arrOutput, index+1, limitNum)
  }

  return arrOutput
}

func getFactor(arr []uint64) uint64{
  factor := uint64(1)
  for i:= range arr{
    factor = factor * arr[i]
    }
  return factor
}

func printArray(arr []uint64){
  for i:= range arr{
    fmt.Printf("[%v]",arr[i])
  }

  fmt.Printf("\n")

  return
}

func main(){
    start := time.Now()

    args := os.Args
    if (len(args) != 3 ) {
      fmt.Println("This is how you use this program:")
      fmt.Println(args[0] + " [number] [number]")
      fmt.Println("Example:")
      fmt.Println(args[0] + " 2 3")
      return
    }

    num_factors,_ := strconv.Atoi(args[1])
    num_len,_ := strconv.Atoi(args[2])

    mod_number := uint64(math.Pow(10,float64(num_len)))

    var factor uint64
    var maxPalindromeFactor uint64

    var v_factors []uint64
    for i:=0;i< num_factors;i++{
      v_factors = append(v_factors,1)
    }

    flgBreak := false

    for ;!flgBreak;{
      i := uint64(0)
      factor = getFactor(v_factors)
      if testPalindrome(factor) && factor > maxPalindromeFactor{
        maxPalindromeFactor = factor
      }
      flgBreak = checkBreak(v_factors, mod_number-1)
      v_factors = modArray(v_factors, i, mod_number)
  }

  fmt.Printf("Duration: %v\n",time.Now().Sub(start))
  fmt.Printf("The largest palindrome made from the product of %v %v-digit numbers is %v\n",num_factors,num_len,maxPalindromeFactor)
}
