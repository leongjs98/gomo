package main

import (
	"fmt"
	"os"
	"time"
)

type arg struct {
	longName    string
	shortName   rune
	description string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	timeArg := arg{
		longName:    "time",
		shortName:   't',
		description: "Set the work time/duration of the task.",
	}
	nameArg := arg{
		longName:    "name",
		shortName:   'n',
		description: "Set the name of the task.",
	}
	descArg := arg{
		longName:    "description",
		shortName:   'd',
		description: "Set the title of the task.",
	}
	logArg := arg{
		longName:    "log",
		shortName:   'l',
		description: "Set the output log file name",
	}
  // TODO: put args into slice, []arg
  // TODO: for loop to check argument (DRY but slower than switch)

	args := map[string]string{
		"duration":    "30m",
		"task":        "Untitled Task",
		"description": "",
		"log":         "./log",
	}

  // TODO: argument error checking
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--" + timeArg.longName, "-" + string(timeArg.shortName):
			args[timeArg.longName] = os.Args[i+1]
		case "--" + nameArg.longName, "-" + string(nameArg.shortName):
			args[nameArg.longName] = os.Args[i+1]
		case "--" + descArg.longName, "-" + string(descArg.shortName):
			args[descArg.longName] = os.Args[i+1]
		case "--" + logArg.longName, "-" + string(logArg.shortName):
			args[logArg.longName] = os.Args[i+1]
		}
	}

	name := args["name"]
	description := args["description"]
	duration, err := time.ParseDuration(args["time"])
	check(err)

	fmt.Println("Name:", name)
	fmt.Println("Description:", description)
	fmt.Printf("Time: %02v:%02v (%v)\n\n", int(duration.Minutes()), int(duration.Seconds())%60, duration)

	// TODO: Use newTicker() to prevent leaking
	c := time.Tick(1 * time.Second)
	startTime := time.Now()

	for _ = range c {
		elapsedTime := time.Since(startTime)
		fmt.Printf("\r%02v:%02v", int(elapsedTime.Minutes()), int(elapsedTime.Seconds())%60)

		if elapsedTime.Seconds() > duration.Seconds() {
			fmt.Println("\n\nTask done!")
      fmt.Println("Task Name:", name)
      fmt.Println("Description:", description)
      fmt.Println("Duration:", duration)
      fmt.Println(startTime.Format("2006/02/01 15:04:05"), "-", time.Now().Format("2006/02/01 15:04:05"))
			return
		}
	}

}
