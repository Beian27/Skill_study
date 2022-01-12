package main
import "fmt"

func main(){
   a := 4
   var ptr *int
   ptr = &a
   fmt.Println(a)
   fmt.Println(*ptr)

}
