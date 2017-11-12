package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type SlaveService int


// master sends it to a waiting slave to give a task
func (m *SlaveService) WakeUp(p1 *bool, p2 *bool) {
	// parameters are useless, but RPC needs them
	hasTask <- true
}


// create a slave server for the master to send a task
// TODO: close it ?
func Serve(port string, done chan bool) error {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0"+port)
	if err != nil {
		return err
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		return err
	}

	service := new(SlaveService)
	rpc.Register(service)
	go rpc.Accept(inbound)
	fmt.Println("RPC server (slave) running on ", addy)
	<-done
	fmt.Println("RPC server (slave) turned off")
	return nil
}
