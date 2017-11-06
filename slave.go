package main

import (
    "os/exec"
)

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

