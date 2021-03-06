package main

import (
	"io/ioutil"
	"os"
	"time"
	"strconv"
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

// The keys are the targets
// We store pointers to rules and not rules directly to be able to update the struct
// See https://stackoverflow.com/a/32751792
type Rules map[string]*Rule

type Task struct {
	Rule          Rule
	RequiredFiles RequiredFiles
}

type Result struct {
	Rule   Rule
	Output []byte
}

type RequiredFiles map[string][]byte

type RulesToParents map[string][]string

type Slave struct {
	Addr string
}

func ReadFile(filename string) ([]byte, error) {
	//TODO: Change this function so that we don't read the entire file in memory but slice by slice
	//TODO: Change the code to send the file so that it is a loop on file slices
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

func Milliseconds(duration time.Duration) string {
	return strconv.FormatInt(duration.Nanoseconds() / 1000000, 10)
}
