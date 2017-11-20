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
    if len(readyRules) == 0 {
        fmt.Println("Adding slave to waiting", (*slave).Addr)
	    waitingSlaves = append(waitingSlaves, slave)
        m.reqMutex.Unlock()
        return
    }

    var rule *Rule
    for k, r := range readyRules {
        rule = r
        delete(readyRules, k)
        break
    }
    // We unlock the mutex here as the following loop may take a while
    m.reqMutex.Unlock()

    requiredFiles := make(RequiredFiles)
    for _, dependency := range rule.Dependencies {
        requiredFiles[dependency], err = ReadFile(path.Join(resultDir, dependency))
        if err != nil {
            fmt.Println("Error reading dependency:", err)
            return
        }
    }
    *reply = Task{Rule: *rule, RequiredFiles: requiredFiles}
	return
}

// The method called by slave when they terminate a task
func (m *MasterService) ReceiveResult(result *Result, end *bool) error {
	*end = false
	WriteFile(path.Join(resultDir, result.Rule.Target), result.Output)

    m.reqMutex.Lock()

	executedRules[result.Rule.Target] = true
	updateParents(result.Rule.Target)

	if result.Rule.Target == firstTarget {
		*end = true
		terminate()
        m.reqMutex.Unlock()
		return nil
	}

	if len(readyRules) == 0 {
        m.reqMutex.Unlock()
        return nil
    }

	// If tasks appeared, wake up all slaves for them to ask for work
    slavesToWakeUp := make([]*Slave, len(waitingSlaves))
    copy(slavesToWakeUp, waitingSlaves)
    waitingSlaves = make([]*Slave, 0)
    // We unlock the mutex here as the following loop may take a while
    m.reqMutex.Unlock()

    for _, slave := range slavesToWakeUp {
        fmt.Println("Waking up", (*slave).Addr)
        slaveClient, _ := rpc.Dial("tcp", (*slave).Addr)
        err := slaveClient.Call("SlaveService.WakeUp", true, nil)
        if err != nil {
            fmt.Println("Error waking up slave:", err)
        }
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
