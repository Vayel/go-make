package main

import (
    "fmt"
	"net"
	"net/rpc"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService int

// use map instead of list as it's more efficient to delete elements from map
var waitingSlaves map[*Slave]bool

// Do not care about the parameter `args`
func (m *MasterService) GiveTask(slave *Slave, reply *Task) error {
    for k, rule := range readyRules {
        // TODO: send dependency files
        *reply = Task{Rule: *rule}
        delete(readyRules, k)
		if _, exists := waitingSlaves[slave]; exists {
			delete(waitingSlaves, slave)
		}
	    return nil
    }
    waitingSlaves[slave] = true
	return nil
}

func (m *MasterService) ReceiveResult(result *Result, reply *bool) error {
	executedRules[result.Rule.Target] = "TODO: generated file"
	updateParents(result.Rule.Target)

    if result.Rule.Target == firstTarget {
        terminate()
        return nil
    }

	// contact waiting slaves if some work appeared
	if len(readyRules) != 0 {
		for slave, _ := range waitingSlaves {
			slaveClient, err := rpc.Dial("tcp", (*slave).Addr)
			if err != nil {
				delete(waitingSlaves, slave)
			}

			err = slaveClient.Call("SlaveService.WakeUp", nil, nil)
			if err != nil {
				delete(waitingSlaves, slave)
			}
		}
	}
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
