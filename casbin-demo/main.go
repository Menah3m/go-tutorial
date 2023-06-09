package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s Can %s %s \n", sub, act, obj)
	} else {
		fmt.Printf("%s Can not %s %s \n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)

	}
	check(e, "admin", "mysql", "write")
	check(e, "developer", "mysql", "write")
	check(e, "kang", "mysql", "write")
	check(e, "min", "mysql", "write")
	check(e, "jie", "mysql", "write")

}
