package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

// RPC（Remote Procedure Call）是远程方法调用的缩写，它可以通过网络调用远程对象的方法。
// Go 标准库net/rpc提供了一个简单、强大且高性能的 RPC 实现。仅需编写很少的代码就能实现 RPC 服务

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error", err)
	}
}
