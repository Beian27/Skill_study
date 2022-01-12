package main
import (
	"fmt"
)
func main(){
	// %d 数字 %s 字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url,stockcode,enddate)
	fmt.Println(target_url)
}
