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

func (m *MasterService) GiveTask(slave *Slave, reply *Task) (err error) {
    for k, rule := range readyRules {
		requiredFiles := make(RequiredFiles)
		for _, dependency := range rule.Dependencies {
			requiredFiles[dependency], err = ReadFile(resultDir + dependency)
			if err != nil {
				return err
			}
		}
		*reply = Task{Rule: *rule, RequiredFiles:requiredFiles}
        delete(readyRules, k)
	    return nil
    }
	waitingSlaves = append(waitingSlaves, slave)
	return nil
}

func (m *MasterService) ReceiveResult(result *Result, reply *bool) error {
	WriteFile(resultDir + result.Rule.Target, result.Output)
    executedRules[result.Rule.Target] = true
    updateParents(result.Rule.Target)

    if result.Rule.Target == firstTarget {
        terminate()
        return nil
    }

	// contact waiting slaves if some work appeared
	if len(readyRules) != 0 {
		for _, slave := range waitingSlaves {
			slaveClient, _ := rpc.Dial("tcp", (*slave).Addr)
			slaveClient.Call("SlaveService.WakeUp", nil, nil)
		}
		waitingSlaves = make([]*Slave, 0);
	}
	return nil
}

func Serve(port string) error {
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
