package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func lookInFile(fileName, phrase string, out chan<- int) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("error: ", err.Error())
		out <- 0
		return
	}
	var countOfOcc int
	buff := bufio.NewReader(file)
	for {
		line, err := buff.ReadString('\n')
		if err == io.EOF {
			countOfOcc += strings.Count(line, phrase)
			break
		} else if err != nil {
			break
		}
		countOfOcc += strings.Count(line, phrase)
	}
	out <- countOfOcc
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Print("need phrase that we look for and name of files")
	}

	out := make(chan int, 2)

	for i := 2; i < len(args); i++ {
		go lookInFile(args[i], args[1], out)
	}
	var sum int
	for i := 2; i < len(args); i++ {
		sum += <-out
	}
	fmt.Printf("Count of of occurance phrase '%s' in files = %d", args[1], sum)
}
