package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"io/ioutil"
	"time"
)


func main() {
	now := time.Now()
	logDir := "/Users/tcrone/lib/newrelic/logs"
	sourceFileLocation := logDir + "/newrelic_agent.log"
	destinationFileLocation := sourceFileLocation + "-" + now.Format("2006-01-02-15:04:05")
	fmt.Printf("Tail New Relic")
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
