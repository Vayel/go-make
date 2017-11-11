package main

import (
	"io/ioutil"
	"os"
)

const fileMode = 0644
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
	RequiredFiles RequiredFiles
	// Also pass the files created by previous commands
	// TODO
}

type Result struct {
    Rule Rule
	Bytes []byte
    // The generated file
    // TODO
}

type RequiredFiles map[string][]byte

type RulesToParents map[string][]string

// TODO: the value type is not a string but a file
type ExecutedRules map[string]string

type Slave struct {
    Todo string
}

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	filebytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return filebytes, nil
}

func WriteFile(filename string, bytes []byte) error {
	err := ioutil.WriteFile(filename, bytes, fileMode)
	return err
}

