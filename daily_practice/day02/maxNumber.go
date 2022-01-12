package main

import "fmt"

func main(){
  var a int = 100
  var b int = 200
  var ret int   
  ret = max(a, b)
  fmt.Println(ret)

}

func max(num_1, num_2 int) int {
   var result int
   if (num_1 > num_2){
	result = num_1
   } else {
	result = num_2	
   }
   return result
}
