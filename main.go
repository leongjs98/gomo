package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const TimeFormat string = "2006-01-02T15:04:05"

type arg struct {
	longName    string
	shortName   rune
	description string
}

func formatTime(d time.Duration) string {
	dStr := fmt.Sprintf("%02vh%02vm%02vs", int(d.Hours()), int(d.Minutes()), int(d.Seconds())%60)
	return dStr
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
	logArg := arg{
		longName:    "log",
		shortName:   'l',
		description: "Set the output log file name",
	}
	// TODO: put args into slice, []arg
	// TODO: for loop to check argument (DRY but slower than switch)

	args := map[string]string{
		"duration": "30m",
		"task":     "Untitled Task",
		"log":      "./gomo.log",
	}

	// TODO: argument error checking
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--" + timeArg.longName, "-" + string(timeArg.shortName):
			args[timeArg.longName] = os.Args[i+1]
		case "--" + nameArg.longName, "-" + string(nameArg.shortName):
			args[nameArg.longName] = os.Args[i+1]
		case "--" + logArg.longName, "-" + string(logArg.shortName):
			args[logArg.longName] = os.Args[i+1]
		}
	}

	name := args["name"]
	duration, err := time.ParseDuration(args["time"])
	logfile := args["log"]
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", name)
	fmt.Printf("Time: %02v:%02v (%v)\n\n", int(duration.Minutes()), int(duration.Seconds())%60, duration)

	// TODO: Use newTicker() to prevent leaking
	ticker := time.Tick(1 * time.Second)
	startTime := time.Now()

	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open the file, %s\n", logfile)
		panic(err)
	}
	defer f.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for sig := range sigChan {
			interruptedDuration := time.Since(startTime)
			logMsg := fmt.Sprintf("%s-%s duration=%s task=\"%s\"\n", startTime.Format(TimeFormat), time.Now().Format(TimeFormat), formatTime(interruptedDuration), name)
			fmt.Printf("\nProgram exited, %v\nAppending log...\n", sig)
			_, err := f.WriteString(logMsg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed writing log file, %s\n", logfile)
			}

			os.Exit(0)
		}
	}()

	for _ = range ticker {
		elapsedTime := time.Since(startTime)
		fmt.Printf("\r%02v:%02v", int(elapsedTime.Minutes()), int(elapsedTime.Seconds())%60)

		if elapsedTime.Seconds() > duration.Seconds() {
			logMsg := fmt.Sprintf("%s-%s duration=%s task=\"%s\"\n", startTime.Format(TimeFormat), time.Now().Format(TimeFormat), formatTime(duration), name)
			_, err = f.WriteString(logMsg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write to the file, %s\n", logfile)
			}

			fmt.Println("\n")
			fmt.Printf("Log appended:\n%s", logMsg)
			fmt.Println("\n\nTake a rest:\ngomo --time \"10m\" --name \"Rest\" \n")

			os.Exit(0)
		}
	}

}
