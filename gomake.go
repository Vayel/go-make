package main

import (
    "fmt"
    "os"
    "path"
    "strings"
    "errors"
    "io"
    "bufio"
)

const CommandPrefix = "\t"
const TargetSuffix = ":"
const (
    HeaderType = iota
    CommandType = iota
    OtherType = iota
)

type Rule struct {
	Target string
	Dependencies []string
	Commands []string
}

// The keys are the targets
// We store pointers to rules and not rules directly to be able to update the struct
// See https://stackoverflow.com/a/32751792
type Rules map[string]*Rule

func help() {
    fmt.Println("Help:")
    fmt.Println("gomake path-to-makefile")
}

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

func parse(f *os.File, rules *Rules) (err error) {
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
    if len(os.Args) < 2 {
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
    err = parse(f, &rules)
    if err != nil {
        fmt.Println("Cannot parse Makefile:", err)
        os.Exit(1)
    }

    printRules(&rules)

    f.Close()
}
