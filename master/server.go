package main

import (
	"fmt"
	"net"
	"net/rpc"
	"path"
    "sync"
)

// The exposed type does not matter, the client only looks at its exported
// methods
type MasterService struct {
    reqMutex sync.Mutex
}

var waitingSlaves []*Slave

// The method called by slaves to ask for work
func (m *MasterService) GiveTask(slave *Slave, reply *Task) (err error) {
    m.reqMutex.Lock()
    defer m.reqMutex.Unlock()

	for k, rule := range readyRules {
		requiredFiles := make(RequiredFiles)
		for _, dependency := range rule.Dependencies {
			requiredFiles[dependency], err = ReadFile(path.Join(resultDir, dependency))
			if err != nil {
				return
			}
		}
		*reply = Task{Rule: *rule, RequiredFiles: requiredFiles}
		delete(readyRules, k)
		return
	}
	waitingSlaves = append(waitingSlaves, slave)
	return
}

// The method called by slave when they terminate a task
func (m *MasterService) ReceiveResult(result *Result, end *bool) error {
    m.reqMutex.Lock()
    defer m.reqMutex.Unlock()

	*end = false
	WriteFile(path.Join(resultDir, result.Rule.Target), result.Output)
	executedRules[result.Rule.Target] = true
	updateParents(result.Rule.Target)

	if result.Rule.Target == firstTarget {
		*end = true
		terminate()
		return nil
	}

	// If tasks appeared, wake up all slaves for them to ask for work
	if len(readyRules) != 0 {
		for _, slave := range waitingSlaves {
			slaveClient, _ := rpc.Dial("tcp", (*slave).Addr)
			slaveClient.Call("SlaveService.WakeUp", true, nil)
		}
		waitingSlaves = make([]*Slave, 0)
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
