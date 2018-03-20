package server

import (
	"net/rpc"
	"net"
	"log"
	"time"

	"BCSD/config"
)

type Service struct {
}

type Processor struct {
}

type SrvInfo struct {
	A, B int
}

type Result struct {
	Quo, Rem int
}

func (t *Processor) RegSrv(args *SrvInfo, SrvInfo *int) error {
	//*reply = args.A * args.B
	return nil
}

func (t *Processor) DelSrv(args *SrvInfo, result *Result) error {
	//if args.B == 0 {
	//	return errors.New("divide by zero")
	//}
	//
	//quo.Quo = args.A / args.B
	//quo.Rem = args.A % args.B
	return nil
}


func (t *Service) Start() {
	processor := new(Processor)
	server := rpc.NewServer()
	server.Register(processor)

	l, e := net.Listen("tcp", cofng.GetConfig().Ip)
	defer l.Close()

	if e != nil {
		log.Fatal("listen error:", e)
		return
	}

	go server.Accept(l)
	log.Println("rpc server started!")

	for {
		time.Sleep(1 * time.Second)
	}
}