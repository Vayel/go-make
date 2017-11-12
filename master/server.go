package main

import (
	"net"
	"net/rpc"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService int
var waitingSlaves []*Slave

// Do not care about the parameter `args`
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
