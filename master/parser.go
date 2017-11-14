package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
    "math"
)

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

func ParseSection(f *os.File, rules *Rules, begin int64, end int64) (err error) {
	rd := bufio.NewReader(f)
	var line, target string
	var dependencies []string

    _, err = f.Seek(begin, 0)
    if err != nil {
        return
    }
    // Reach the end of the line
    _, err = rd.ReadString('\n')
    if err != nil {
        return
    }
    // Reach the beginning of the next rule
    for err == nil{
		line, err = rd.ReadString('\n')
		if getLineType(line) == HeaderType {
            break
        }
    }
    
    for pos, e := f.Seek(0, 1); e == nil && err == nil && pos <= end; {
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
		line, err = rd.ReadString('\n')
	}
	if err != io.EOF {
		return
	}
    return nil
}


func Parse(f *os.File, rules *Rules) (err error) {
    fi, err := f.Stat()
    if err != nil {
        return
    }
    size := fi.Size()
    middle := int64(math.Floor(float64(size)/2))
    ParseSection(f, rules, 0, middle)
    ParseSection(f, rules, middle, size)
	return
}
