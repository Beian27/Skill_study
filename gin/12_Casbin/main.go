package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, err := xormadapter.NewAdapter("mysql", "root:Zhy5224146.+@tcp(127.0.0.1:3306)/casbin", true)
	e, err := casbin.NewEnforcer("./12_Casbin/model.conf", a)

	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read" // 用户对资源执行的操作。
	//addPolicy, err := e.AddPolicy("alice", "data2", "read")
	//fmt.Println(addPolicy)
	//fmt.Println(err)
	//policy, err := e.AddGroupingPolicy("alice", "data2_admin")
	//fmt.Println(policy)
	//fmt.Println(err)
	//e.UpdatePolicy([]string{"alice", "data2", "read"},[]string{"alice", "data3", "read"})
	//filteredPolicy := e.GetFilteredPolicy(0, "alice")
	//fmt.Println(filteredPolicy)

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// 处理err
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("pass")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("拒绝请求，抛出异常")
	}


}
