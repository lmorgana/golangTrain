package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type fileParam struct {
	fileName string
	CountOfL int   //newline counts
	CountOfC int   //byte counts
	CountOfW int   //word counts
	CountOfM int   //character counts
	err      error //if something went wrong
}

type flags struct {
	l *bool
	c *bool
	m *bool
	w *bool
}

func (f *fileParam) init(fileName string) {
	f.CountOfL = 0
	f.CountOfC = 0
	f.CountOfW = 0
	f.CountOfM = 0
	f.fileName = fileName
}

func (f *fileParam) sum(other *fileParam) {
	f.CountOfL += other.CountOfL
	f.CountOfC += other.CountOfC
	f.CountOfW += other.CountOfW
	f.CountOfM += other.CountOfM
}

func printParam(pm *fileParam, fl *flags) {
	if pm.err != nil {
		fmt.Println("wc:", pm.err.Error())
	} else {
		if *fl.l {
			fmt.Printf("%8d", pm.CountOfL)
		} else if *fl.w {
			fmt.Printf("%8d", pm.CountOfW)
		} else if *fl.c {
			fmt.Printf("%8d", pm.CountOfC)
		} else if *fl.m {
			fmt.Printf("%8d", pm.CountOfM)
		} else {
			fmt.Printf("%8d%8d%8d %8s\n", pm.CountOfL, pm.CountOfW, pm.CountOfC)
		}
		fmt.Printf(" %8s\n", pm.fileName)
	}
}

func wcString(line string, buffPm *fileParam) {
	if strings.Contains(line, "\n") {
		buffPm.CountOfL++
	}
	buffPm.CountOfW += len(strings.Fields(line))
	buffPm.CountOfM += utf8.RuneCountInString(line)
	for _, val := range line {
		buffPm.CountOfC += utf8.RuneLen(val)
	}
}

func wcFile(fileName string, buffPm *fileParam) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err == nil {
		buff := bufio.NewReader(file)
		for {
			line, err := buff.ReadString('\n')
			if err == io.EOF {
				wcString(line, buffPm)
				break
			}
			if err != nil {
				buffPm.err = err
				break
			}
			wcString(line, buffPm)
		}
	}
	buffPm.err = err
}

func wcForFiles(namesOfFiles []string, fl *flags) {
	var buffPm *fileParam //buffer for count a file parameter
	buffPm = new(fileParam)

	var totalPm *fileParam //count for total count of parameters
	totalPm = new(fileParam)
	totalPm.init("total")

	for i := 0; i < len(namesOfFiles); i++ {
		buffPm.init(namesOfFiles[i])
		wcFile(namesOfFiles[i], buffPm)
		printParam(buffPm, fl)
		totalPm.sum(buffPm)
	}
	if len(namesOfFiles) > 0 {
		printParam(totalPm, fl)
	}
}

func wcForStdin(fl *flags) {
	var err error
	var buffPm *fileParam
	buffPm = new(fileParam)
	buffPm.init("")
	buff := bufio.NewReader(os.Stdin)
	for {
		line, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		wcString(line, buffPm)
	}
	buffPm.err = err
	printParam(buffPm, fl)
}

func main() {
	args := os.Args
	fl := new(flags)
	fl.l = flag.Bool("l", false, "-l")
	fl.c = flag.Bool("c", false, "-c")
	fl.m = flag.Bool("m", false, "-m")
	fl.w = flag.Bool("w", false, "-w")
	flag.Parse()
	if len(args) > 1 {
		wcForFiles(flag.Args(), fl)
	} else {
		wcForStdin(fl)
	}
	return
}
