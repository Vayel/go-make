package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
    "os/exec"
)

var done chan bool
var hasTask chan bool
var task Task


func work(task Task) (err error) {
    fmt.Println("Working on ", task.Rule.Target)
    for _, cmd := range task.Rule.Commands {
        if e := exec.Command("sh", "-c", cmd).Run(); e != nil {
            return e
        }
    }
    return
}

func help() {
	fmt.Println("Help:")
	fmt.Println("\tslave master-rpc-addr master-rpc-port")
	fmt.Println("\nExample:")
	fmt.Println("\tslave localhost 10000")
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments")
		help()
		os.Exit(1)
	}

	masterAddr := os.Args[1]
	masterPort := os.Args[2]
	slaveAddr := os.Args[3]
	slavePort := os.Args[4]
	client, err := rpc.Dial("tcp", masterAddr+":"+masterPort)
	if err != nil {
		log.Fatal(err)
	}

	task = Task{}
	slave := Slave{Addr:slaveAddr+":"+slavePort}
	var result Result
    var reply bool

	// start RPC server for the master to contact us
	// if no task available when calling GiveTask
	// (it's more efficient than starting it and closing it
	// dynamically when needed)
	done = make(chan bool, 1)
	go Serve(slavePort)
	// TODO: handle errors

	/*
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
	*/


    for {
        err = client.Call("MasterService.GiveTask", &slave, &task)
        if err != nil {
            fmt.Println(err)
            return
        }

        if len(task.Rule.Target) == 0 {
			<-hasTask
			continue
        }

        work(task)
        result = Result{Rule: task.Rule}
        err = client.Call("MasterService.ReceiveResult", &result, &reply)
        if err != nil {
            fmt.Println(err)
            return
        }

	    task = Task{}
    }
}
