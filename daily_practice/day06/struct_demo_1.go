package main

import "fmt"

type Books struct {
   title string
   author string 
   subject string
   book_id int
}

func main(){
  var book1 Books
  
  book1.title = "Go 语言"
  book1.author = "Beian27"
  book1.subject = "工学"
  book1.book_id = 11111111

  fmt.Println(book1)
}
