package main
import "fmt"
func main(){
	sum := 1
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	
	for sum <= 10{
		sum += sum
	}
	fmt.Println(sum)
}
