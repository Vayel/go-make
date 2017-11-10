package main

import (
    "fmt"
	"net"
	"net/rpc"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService int
var waitingSlaves []*Slave

// Do not care about the parameter `args`
func (m *MasterService) GiveTask(slave *Slave, reply *Task) error {
    for k, rule := range readyRules {
        // TODO: send dependency files
        *reply = Task{Rule: *rule}
        delete(readyRules, k)
	    return nil
    }
    waitingSlaves = append(waitingSlaves, slave)
	return nil
}

func (m *MasterService) ReceiveResult(result *Result, reply *bool) error {
    // TODO: save generated file
    executedRules[result.Rule.Target] = "generated file"
    updateParents(result.Rule.Target)

    if result.Rule.Target == firstTarget {
        terminate()
        return nil
    }

    // TODO: contact waiting slaves if some work appeared

	return nil
}

func Serve(port string, done chan bool) error {
    addr := "0.0.0.0:" + port
	addy, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		return err
	}

	service := new(MasterService)
	rpc.Register(service)
	go rpc.Accept(inbound)
    fmt.Println("RPC server running on", addr)
    <-done
    fmt.Println("RPC server turned off")
	return nil
}
