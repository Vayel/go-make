package main

import (
    "fmt"
    "os"
    "os/exec"
)

func help() {
	fmt.Println("Help:")
	fmt.Println("\tsequential path-to-makefile rule-to-execute")
	fmt.Println("\nExamples:")
	fmt.Println("\tsequential Makefile all")
	fmt.Println("\tsequential ../MyMakefile test.c")
}

func getDependentTargets(rule *Rule, rules *Rules) (dependencies []*Rule) {
    for _, dep := range rule.Dependencies {
        if r, isPresent := (*rules)[dep]; isPresent { // The dependency is a target itself
            dependencies = append(dependencies, r)
        }
    }
    return
}

func execute(target string, rules *Rules) (err error) {
    rule := (*rules)[target]
    dependencies := getDependentTargets(rule, rules)

    for _, dep := range dependencies {
        execute(dep.Target, rules)
    }

    for _, cmd := range rule.Commands {
        if e := exec.Command("sh", "-c", cmd).Run(); e != nil {
            return e
        }
    }

    return
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Not enough arguments")
        help()
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
        os.Exit(1)
    }
	defer f.Close()

    rules := make(Rules)
    err = Parse(f, &rules)
    if err != nil {
        fmt.Println("Cannot parse Makefile:", err)
        os.Exit(1)
    }

	firstTarget := os.Args[2]
	if _, present := rules[firstTarget]; !present {
		fmt.Printf("Invalid target '%s'\n", firstTarget)
		os.Exit(1)
	}

    if e := execute(firstTarget, &rules); e != nil {
        fmt.Printf("Error executing target '%s': %s\n", firstTarget, e)
        os.Exit(1)
    }
}
