package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Usage: btail <path-to-file>")
		return
	}
	now := time.Now()
	sourceFileLocation := args[1]
	destinationFileLocation := sourceFileLocation + "-" + now.Format("2006-01-02-15:04:05")
	fmt.Printf("Backup and Tail [%s]\n", sourceFileLocation)
	input, err := ioutil.ReadFile(sourceFileLocation)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(destinationFileLocation, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFileLocation)
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(sourceFileLocation, nil, 0644)
	if err != nil {
		fmt.Println("Error creating new log file", sourceFileLocation)
		fmt.Println(err)
		return
	}
	t, err := tail.TailFile(sourceFileLocation, tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
