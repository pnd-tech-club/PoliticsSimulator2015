package main

import (
  "fmt"
)

/* used on stats that are utilize the -5 to 5 range
- This is a variadic function */

/* takes a []int as input in the form: comparestats(thisvar...)*/
func comparestats(nums ...int) { // <WIP>
  fmt.Print("DEBUG: ", nums)
  n := 0
  var newnums []int

  for _, num := range nums{
    n+=1
    newnums[n] += num
  }
  fmt.Println(newnums)

}
