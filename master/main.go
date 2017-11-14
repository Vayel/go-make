package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"path"
	"time"
)

// The keys are the targets
// We store pointers to rules and not rules directly to be able to update the struct
// See https://stackoverflow.com/a/32751792
type Rules map[string]*Rule

// Use a map to efficiently determine if a rule has been executed
type ExecutedRules map[string]bool

var rules Rules
var rulesToParents RulesToParents
var readyRules Rules // Use a map to avoid duplicates
var executedRules ExecutedRules
var firstTarget string
var done chan bool
var resultDir string

func help() {
	fmt.Println("Help:")
	fmt.Println("\tmaster path-to-makefile rule-to-execute rpc-port result-dir")
	fmt.Println("\nExamples:")
	fmt.Println("\tmaster Makefile all 10000 outputfiles/")
	fmt.Println("\tmaster ../MyMakefile test.c 10000 dir/outputfiles/")
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
		slaveClient, _ := rpc.Dial("tcp", (*slave).Addr)
		slaveClient.Call("SlaveService.ShutDown", true, nil)
	}

	// Kill the RPC server
	done <- true
}

func getAbsolutePath(relPath string) (string, error) {
	wdir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(wdir, relPath), nil
}

func main() {
	// Time measures
	logfile, errf := os.OpenFile("time_master.log", os.O_WRONLY|os.O_CREATE, 0644)
	if errf != nil {
		log.Fatal(errf)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	startTime := time.Now()
	if len(os.Args) != 5 {
		fmt.Println("Invalid number of arguments")
		help()
		os.Exit(1)
	}

	resultDir = os.Args[4]
	if stat, err := os.Stat(resultDir); err != nil || !stat.IsDir() {
		fmt.Println("Result directory does not exist: " + resultDir)
		os.Exit(1)
	}

	path := os.Args[1]
	path, err := getAbsolutePath(path)
	if err != nil {
		fmt.Println("Cannot open Makefile:", err)
		os.Exit(1)
	}

	f, err := os.Open(path)
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

	elapsedTime := time.Since(startTime)
	log.Println("Time (master) : ", elapsedTime)
}
