package main

import (
    "fmt"
    "os"
    "path"
    "strings"
    "os/exec"
)

func help() {
    fmt.Println("Help:")
    fmt.Println("\tgomake path-to-makefile rule-to-execute")
    fmt.Println("\nExamples:")
    fmt.Println("\tgomake Makefile all")
    fmt.Println("\tgomake ../MyMakefile test.c")
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

    rules := make(Rules)
    err = Parse(f, &rules)
    if err != nil {
        fmt.Println("Cannot parse Makefile:", err)
        os.Exit(1)
    }

    target := os.Args[2]
    if _, isPresent := rules[target]; !isPresent {
        fmt.Printf("The rule '%s' does not appear in the Makefile\n", target)
        os.Exit(1)
    }

    if e := execute(target, &rules); e != nil {
        fmt.Printf("Error executing target '%s': %s\n", target, e)
        os.Exit(1)
    }

    f.Close()
}
