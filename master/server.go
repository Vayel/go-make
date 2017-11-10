package main

import (
	"net"
	"net/rpc"
	"io/ioutil"
	"os"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService int
var waitingSlaves []*Slave

// Do not care about the parameter `args`
func (m *MasterService) GiveTask(slave *Slave, reply *Task) error {
    for k, rule := range readyRules {
        *reply = Task{Rule: *rule}
        delete(readyRules, k)
	    return nil
    }
    waitingSlaves = append(waitingSlaves, slave)
	return nil
}

// We give the required file to the slave
func (m *MasterService) GiveFile(slave *Slave, reply *([]byte)) error {
	filename := "makefiles/1" // Make this a parameter
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	*reply, err = ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return nil
}

func (m *MasterService) ReceiveResult(result *Result, reply *bool) error {
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
