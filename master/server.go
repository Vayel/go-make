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
		*reply = Task{Rule: *rule}
		delete(readyRules, k)
		return nil
	}
	waitingSlaves[slave] = true
	return nil
}

func (m *MasterService) ReceiveResult(result *Result, reply *bool) error {
	executedRules[result.Rule.Target] = "TODO: generated file"
	updateParents(result.Rule.Target)

	// contact waiting slaves if some work appeared

	var toRemove []*Slave // temporary list of slaves to which we gave a task
	// needed because we can't delete elements from a list while looping
	for slave, _ := range waitingSlaves {
		if len(readyRules) != 0 {
			// get a random task from the ReadyRules list
			var rule *Rule
			for _, r := range readyRules {
				rule = r
				break
			}

			task := Task{Rule: *rule}

			slaveClient, err := rpc.Dial("tcp", (*slave).Addr+":40000")
			if err != nil {
				return err
			}
			err = slaveClient.Call("MasterService.ReceiveTask", &task, nil)
			if err != nil {
				return err
			}
			toRemove = append(toRemove, slave)
		} else {
			break
		}
	}

	// delete slaves which have now a task
	for _,slave := range toRemove {
		delete(waitingSlaves, slave)
	}

	return nil
}


// master sends it to a waiting slave to give it a task
func (m *MasterService) ReceiveTask(task *Task, reply* Result) {
	// TODO: run task

	// work(task)
	result := Result{Rule: task.Rule}

	// err := ???.Call("MasterService.ReceiveResult", &result, &reply)
	if err != nil {
		fmt.Println(err)
	}
	return
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
