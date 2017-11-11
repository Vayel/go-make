package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
    "os/exec"
)

var filedir string = "outputfiles/"

func writeFiles(requiredFiles RequiredFiles) error {
	for filename, bytes := range requiredFiles {
		err := WriteFile(filedir + filename, bytes)
		if(err != nil) {
			return err
		}
	}
	return nil
}

func work(task Task) (err error) {
	writeFiles(task.RequiredFiles)
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
	slave := Slave{Todo: "todo"}
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
		fileResult, err := ReadFile(filedir + task.Rule.Target)
		if err != nil {
            fmt.Println(err)
			return
		}
		result = Result{Rule: task.Rule, Bytes: fileResult}
        err = client.Call("MasterService.ReceiveResult", &result, &reply)
        if err != nil {
            fmt.Println(err)
            return
        }

	    task = Task{}
    }
}
