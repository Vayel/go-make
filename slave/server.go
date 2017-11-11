package main

import (
	"net"
	"net/rpc"
)

type SlaveService int


// master sends it to a waiting slave to give a task
func (m *SlaveService) ReceiveTask(toDo *Task, reply* Result) {
	task = *toDo
	hasTask <- true
}


// create a slave server for the master to send a task
// TODO: close it ?
func Serve(port string) error {
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
	rpc.Accept(inbound)
	return nil
}
