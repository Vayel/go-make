package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type SlaveService int

// master sends it to a waiting slave to tell him tasks are available
func (m *SlaveService) WakeUp(p1 bool, p2 *bool) error {
	// parameters are useless, but RPC needs them
    fmt.Println("SlaveService.WakeUp received")
	hasTask <- true
	return nil
}

// master sends it to a waiting slave to tell him to shut down
func (m *SlaveService) ShutDown(p1 bool, p2 *bool) error {
    fmt.Println("SlaveService.ShutDown received")
	hasTask <- false
	done <- true
	return nil
}

func createServer(port string) (*net.TCPListener, error) {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+port)
	if err != nil {
		return nil, err
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		return nil, err
	}
	return inbound, nil
}

// create a slave server for the master to send a task
func Serve(inbound *net.TCPListener) {
	service := new(SlaveService)
	rpc.Register(service)
	go rpc.Accept(inbound)
	<-done
	fmt.Println("RPC server (slave) turned off")
}
