package main

import (
	"net"
	"net/rpc"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService int
var waitingSlaves []*Slave
var slavefiledir string = "outputfiles/"

// Do not care about the parameter `args`
func (m *MasterService) GiveTask(slave *Slave, reply *Task) error {
    for k, rule := range readyRules {
		var requiredFiles RequiredFiles = make(RequiredFiles)
		var err error
		for _, dependency := range rule.Dependencies {
			requiredFiles[dependency], err = ReadFile(slavefiledir + dependency)
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
	WriteFile(slavefiledir + result.Rule.Target, result.Bytes)
    executedRules[result.Rule.Target] = "TODO: generated file"
    updateParents(result.Rule.Target)

    // TODO: contact waiting slaves if some work appeared

	return nil
}

func Serve(port string) error {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+port)
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
