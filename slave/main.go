package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
    "os/exec"
)

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
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments")
		help()
		os.Exit(1)
	}

	addr := os.Args[1]
	port := os.Args[2]
	client, err := rpc.Dial("tcp", addr+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	task := Task{}
	slave := Slave{Todo: "todo", Addr:"0.0.0.0"} // TODO: get addr
	var result Result
    var reply bool
    for {
        err = client.Call("MasterService.GiveTask", &slave, &task)
        if err != nil {
            fmt.Println(err)
            return
        }
        if len(task.Rule.Target) == 0 {
            break
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

    // TODO: start RPC server for the master to contact us
	// TODO: when task given, close server and back to infinite loop ?

	// if no task given by the master, open a server and wait for a task
	err = Serve(40000) // TODO: choose port ?
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
}
