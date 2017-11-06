package main

import (
    "fmt"
    "net"
    "net/rpc"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService int

// Do not care about the parameter `args`
func (m *MasterService) GiveTask(args *int, reply *Task) error {
    // TODO
    fmt.Println("GiveTask called")
    *reply = Task{}
    return nil
}

func (m *MasterService) ReceiveResult(args *Result, reply *bool) error {
    // TODO
    return nil
}

func Serve(port string) error {
    addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:" + port)
	if err != nil {
        return err
    }

    inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
        return err
    }

    service := new(MasterService)
    rpc.Register(service)
    rpc.Accept(inbound)
    return nil
}
