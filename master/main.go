package main

import (
	"fmt"
	"net/rpc"
	"os"
	"time"
)

// Use a map to efficiently determine if a rule has been executed
type ExecutedRules map[string]bool

var rules Rules
var rulesToParents RulesToParents
var readyRules Rules // Use a map to avoid duplicates
var executedRules ExecutedRules
var firstTarget string
var done chan bool

func help() {
	fmt.Println("Help:")
	fmt.Println("\tmaster path-to-makefile rule-to-execute rpc-port log-path")
	fmt.Println("\nExamples:")
	fmt.Println("\tmaster Makefile all 10000 ~/logdir/")
	fmt.Println("\tmaster ../MyMakefile test.c 10000 /tmp/logdir/master.json")
}

// With this function, we can easily access the parents of a given rule, that is
// the rules depending on that rule
func linkRulesToParents(rules *Rules, parent string, mapping *RulesToParents) {
	rule := (*rules)[parent]
	for _, dep := range rule.Dependencies {
		if (*mapping)[dep] == nil {
			(*mapping)[dep] = make([]string, 0)
		}
		(*mapping)[dep] = append((*mapping)[dep], parent)
		linkRulesToParents(rules, dep, mapping)
	}
	if len(rule.Dependencies) == 0 {
		readyRules[parent] = rule
	}
}

// When a dependency has been executed, some of its parents may become ready
func updateParents(child string) {
	for _, parent := range rulesToParents[child] {
		if isReady(rules[parent]) {
			readyRules[parent] = rules[parent]
		}
	}
}

// A rule is ready if it has not been executed and if all its dependencies have
// been executed
func isReady(rule *Rule) bool {
	if _, present := executedRules[rule.Target]; present {
		return false
	}
	for _, dep := range rule.Dependencies {
		if _, present := executedRules[dep]; !present {
			return false
		}
	}
	return true
}

func terminate() {
	fmt.Println(firstTarget, "rule has been computed!")

	// Tell waiting slaves to shutdown
	for _, slave := range waitingSlaves {
        fmt.Println("Shuting down", (*slave).Addr)
		slaveClient, _ := rpc.Dial("tcp", (*slave).Addr)
		err := slaveClient.Call("SlaveService.ShutDown", true, nil)
        if err != nil {
            fmt.Println("Error shuting down slave:", err)
        }
	}

	// Kill the RPC server
	done <- true
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Invalid number of arguments:", os.Args)
		help()
		os.Exit(1)
	}

	startTime := time.Now()

	path_ := os.Args[1]
	f, err := os.Open(path_)
	if err != nil {
		fmt.Println("Cannot open Makefile:", err)
		f.Close()
		os.Exit(1)
	}
	defer f.Close()

	rules = make(Rules)
	err = Parse(f, &rules)
	if err != nil {
		fmt.Println("Cannot parse Makefile:", err)
		f.Close()
		os.Exit(1)
	}

	firstTarget = os.Args[2]
	if _, present := rules[firstTarget]; !present {
		fmt.Printf("Invalid target '%s'\n", firstTarget)
		os.Exit(1)
	}

	executedRules = make(ExecutedRules)
	readyRules = make(Rules)
	rulesToParents = make(RulesToParents)
	linkRulesToParents(&rules, firstTarget, &rulesToParents)

	port := os.Args[3]
	done = make(chan bool, 1)
	err = Serve(port)
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}

	logfile, errf := os.Create(os.Args[4])
	if errf != nil {
		panic(errf)
	}
	defer logfile.Close()
	elapsedTime := time.Since(startTime)
	fmt.Fprintf(logfile, "{\"total\": " + Milliseconds(elapsedTime) + "}")
}
