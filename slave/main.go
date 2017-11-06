package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

/*
func getDependentTargets(rule *Rule, rules *Rules) (dependencies []*Rule) {
    for _, dep := range rule.Dependencies {
        if r, isPresent := (*rules)[dep]; isPresent { // The dependency is a target itself
            dependencies = append(dependencies, r)
        }
    }
    return
}

func Execute(target string, rules *Rules) (err error) {
    rule := (*rules)[target]
    dependencies := getDependentTargets(rule, rules)

    for _, dep := range dependencies {
        Execute(dep.Target, rules)
    }

    for _, cmd := range rule.Commands {
        if e := exec.Command("sh", "-c", cmd).Run(); e != nil {
            return e
        }
    }

    return
}
*/

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
	client.Call("MasterService.GiveTask", 0, &task)

	/*
	   var reply bool
	   for task != nil {
	       result := work(task)
	       client.Call("MasterService.ReceiveResult", &result, &reply)
	       client.Call("MasterService.GiveTask", 0, &task)
	   }
	*/
}
