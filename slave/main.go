package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
    "os/exec"
)

func work(task Task) (err error) {
    fmt.Println("work on ", task.Rule)
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
	var result Result
    var reply bool
    for {
        err = client.Call("MasterService.GiveTask", 0, &task)
        if err != nil {
            fmt.Println(err)
            break
        }
        if len(task.Rule.Target) == 0 {
            continue
        }

        work(task)
        result = Result{Rule: task.Rule}
        err = client.Call("MasterService.ReceiveResult", &result, &reply)
        if err != nil {
            fmt.Println(err)
            break
        }

	    task = Task{}
    }
}
