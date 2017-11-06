package main

import (
    "os"
    "fmt"
    "strings"
    "errors"
    "io"
    "bufio"
)

// The keys are the targets
// We store pointers to rules and not rules directly to be able to update the struct
// See https://stackoverflow.com/a/32751792
type Rules map[string]*Rule

func getLineType(line string) int {
    if strings.HasPrefix(line, CommandPrefix) {
        return CommandType
    } else if len(strings.Fields(line)) > 0 { // If the line has not blanks only
        return HeaderType
    }
    return OtherType
}

func parseHeader(header string) (target string, dependencies []string, err error) {
    splitHeader := strings.Split(header, TargetSuffix)
    if len(splitHeader) != 2 {
        err = errors.New(fmt.Sprintf("The header '%s' is not the form of 'target: deps'", header))
        return
    }

    target = strings.Trim(splitHeader[0], " ")
    dependencies = strings.Fields(splitHeader[1]) // Does not work if a dependency contains spaces
    return
}

func parseCommand(line string) string {
    return strings.TrimSpace(line)
}

func Parse(f *os.File, rules *Rules) (err error) {
    rd := bufio.NewReader(f)
    var line, target string
    var dependencies []string

    for err == nil {
        line, err = rd.ReadString('\n')

        if getLineType(line) == HeaderType {
            target, dependencies, err = parseHeader(line)
            if err != nil {
                return err
            }

            (*rules)[target] = &Rule{target, dependencies, make([]string, 0)}
        } else if getLineType(line) == CommandType {
            cmd := parseCommand(line)
            (*rules)[target].Commands = append((*rules)[target].Commands, cmd)
        }
    }

    if err != io.EOF {
        return err
    }

    return nil
}
