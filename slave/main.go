package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"path"
	"time"
)

var dependencyDir string

var done chan bool
var hasTask chan bool
var task Task

func writeFiles(requiredFiles RequiredFiles) error {
	for filename, bytes := range requiredFiles {
		err := WriteFile(path.Join(dependencyDir, filename), bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func work(task Task) (err error) {
	fmt.Println("Begin", task.Rule.Target)
	writeFiles(task.RequiredFiles)
	for _, cmd := range task.Rule.Commands {
		if e := exec.Command("sh", "-c", cmd).Run(); e != nil {
			return e
		}
	}
	fmt.Println("End", task.Rule.Target, "\n")
	return
}

func help() {
	fmt.Println("Help:")
	fmt.Println("\tslave master-rpc-addr master-rpc-port slave-rpc-addr slave-rpc-port dependency-dir")
	fmt.Println("\nExample:")
	fmt.Println("\tslave 129.6.12.82 10000 129.6.12.81 40000 ~/rpc-go/")
}

func main() {
	startTime := time.Now()
	var workTime, waitTime time.Duration = 0, 0
	hasTask = make(chan bool, 1)

	if len(os.Args) < 6 {
		fmt.Println("Not enough arguments")
		help()
		os.Exit(1)
	}

	masterAddr := os.Args[1]
	masterPort := os.Args[2]
	slaveAddr := os.Args[3]
	slavePort := os.Args[4]
	dependencyDir = os.Args[5]

	if stat, err := os.Stat(dependencyDir); err != nil || !stat.IsDir() {
		fmt.Println("Not a directory: " + dependencyDir)
		os.Exit(1)
	}

	client, err := rpc.Dial("tcp", masterAddr+":"+masterPort)
	if err != nil {
		log.Fatal(err)
	}

	task = Task{}
	slave := Slave{Addr: slaveAddr + ":" + slavePort}
	var result Result
	var end bool

	// start RPC server for the master to contact us
	// if no task available when calling GiveTask
	// (it's more efficient than starting it and closing it
	// dynamically when needed)
	done = make(chan bool, 1)
	inbound, err := createServer(slavePort)
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
	fmt.Println("RPC server (slave) running on ", slaveAddr+":"+slavePort)
	go Serve(inbound)

	for {
		err = client.Call("MasterService.GiveTask", &slave, &task)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(task.Rule.Target) == 0 {
			fmt.Println("Wait for task\n")
			startWaitTime := time.Now()
			running := <-hasTask
			waitTime += time.Since(startWaitTime)
			if !running {
				break
			}
			continue
		}

		startWorkTime := time.Now()
		work(task)
		workTime += time.Since(startWorkTime)

		fileResult, err := ReadFile(dependencyDir + task.Rule.Target)
		if err != nil {
			fmt.Println(err)
			return
		}
		result = Result{Rule: task.Rule, Output: fileResult}
		err = client.Call("MasterService.ReceiveResult", &result, &end)
		if err != nil {
			fmt.Println(err)
			return
		}
		if end {
			break
		}

		task = Task{}
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Time (slave) : ", elapsedTime)
	fmt.Println("Work Time (slave) : ", workTime)
	fmt.Println("Wait Time (slave) : ", waitTime)
}
