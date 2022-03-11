package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

// casbin是一个强大、高效的访问控制库。支持常用的多种访问控制模型，如ACL/RBAC/ABAC等。可以实现灵活的访问权限控制
func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforcer failed:%v\n", err)
	}

	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
}
