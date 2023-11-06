package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

const TimeFormat string = "2006-01-02T15:04:05"

type flagStruct struct {
	name  string
	value string
	usage string
}

var durationStruct = flagStruct{
	name:  "duration",
	value: "30m",
	usage: "Set the duration of the task.",
}

var taskStruct = flagStruct{
	name:  "task",
	value: "Untitled task",
	usage: "Set the name and/or description of the task.",
}

var logStruct = flagStruct{
	name:  "log",
	value: "./gomo.log",
	usage: "Set path to the output log file.",
}

func formatTime(d time.Duration) string {
	dStr := fmt.Sprintf("%02vh%02vm%02vs", int(d.Hours()), int(d.Minutes())%60, int(d.Seconds())%60)
	return dStr
}

func main() {

	var durationFlag, taskFlag, logFlag string
	flag.StringVar(&durationFlag, durationStruct.name, durationStruct.value, durationStruct.usage)
	flag.StringVar(&taskFlag, taskStruct.name, taskStruct.value, taskStruct.usage)
	flag.StringVar(&logFlag, logStruct.name, logStruct.value, logStruct.usage)

	flag.Parse()
	fmt.Println(durationFlag, taskFlag, logFlag)

	duration, err := time.ParseDuration(durationFlag)
	logfile := logFlag
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", taskFlag)
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
			logMsg := fmt.Sprintf("%s-%s duration=%s task=\"%s\"\n", startTime.Format(TimeFormat), time.Now().Format(TimeFormat), formatTime(interruptedDuration), taskFlag)
			fmt.Printf("\nProgram exited, %v\nAppending log...\n", sig)
			_, err := f.WriteString(logMsg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed writing log file, %s\n", logfile)
			}

			os.Exit(1)
		}
	}()

	for _ = range ticker {
		elapsedTime := time.Since(startTime)
		fmt.Printf("\r%02v:%02v", int(elapsedTime.Minutes()), int(elapsedTime.Seconds())%60)

		if elapsedTime.Seconds() > duration.Seconds() {
			logMsg := fmt.Sprintf("%s-%s duration=%s task=\"%s\"\n", startTime.Format(TimeFormat), time.Now().Format(TimeFormat), formatTime(duration), taskFlag)
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
