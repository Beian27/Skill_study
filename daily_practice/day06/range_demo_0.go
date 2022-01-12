package main

import "fmt"

func main(){
  nums := []int{1,2,3,4}
  sum := 0 
  
  for index, num := range nums {
    fmt.Println(index)
    sum += num 
  }
  fmt.Printf("Sum = %d\n",sum)

  kvs := map[string]string{"a":"apple","b":"banna"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n",k,v)
  }
  
  for i, c := range "Beian"{
    fmt.Printf("%d -> %s",i,c)
  }
}
