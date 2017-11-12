package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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
	fmt.Println("\tmaster Makefile all 10000")
	fmt.Println("\tmaster ../MyMakefile test.c 10000")
}

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

func updateParents(child string) {
    for _, parent := range rulesToParents[child] {
        if isReady(rules[parent]) {
		    readyRules[parent] = rules[parent]
        }
    }
}

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
    // Kill the RPC server
    done <- true
    // TODO:
    //    * kill waiting slaves
}

func getAbsolutePath(relPath string) (string, error) {
	wdir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(wdir, relPath), nil
}

func printRules(rules *Rules) {
	for target, rule := range *rules {
		fmt.Print(target, ": ", strings.Join(rule.Dependencies, " "), "\n")
		for _, cmd := range rule.Commands {
			fmt.Println(CommandPrefix, cmd)
		}
	}
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Invalid number of arguments")
		help()
		os.Exit(1)
	}

	if stat, err := os.Stat(os.Args[4]); err != nil || !stat.IsDir() {
		fmt.Println("Not a directory: " + os.Args[4])
		os.Exit(1)
	}
	resultDir = os.Args[4]

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
	err = Serve(port, done)
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
}
