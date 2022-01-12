package main

var x, y int

var ( //声明全局变量
   a int
   b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
    //这种不带声明格式的只能在函数体中出现
    g, h := 123, "ZhangHongYu"
    println(a,b,c,d,e,f,g,h)
}
