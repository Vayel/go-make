package main

const CommandPrefix = "\t"
const TargetSuffix = ":"
const (
	HeaderType  = iota
	CommandType = iota
	OtherType   = iota
)

type Rule struct {
	Target       string
	Dependencies []string
	Commands     []string
}

type Task struct {
	Rule Rule
	// Also pass the files created by previous commands
	// TODO
}

type Result struct {
    Rule Rule
    // The generated file
    // TODO
}

type RulesToParents map[string][]string
