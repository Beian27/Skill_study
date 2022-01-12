package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main(){
  fmt.Println(Books{"Go 语言", "大卫·约翰逊", "工科", 101024})
}
