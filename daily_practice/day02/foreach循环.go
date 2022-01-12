package main
import "fmt"

func main(){
	strings := []string{"google","runoob"}
	for k,v := range strings {
		data := fmt.Sprintf("Key:%d  Value:%s",k,v)
		fmt.Println(data)
	}
}
