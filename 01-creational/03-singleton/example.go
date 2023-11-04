package main

import "fmt"

func main() {
	//log := getLoggerInstance()
	//
	//log.SetLogLevel(1)
	//log.Log("This is a log message")
	//
	//log = getLoggerInstance()
	//log.SetLogLevel(2)
	//log.Log("This is a log message")
	//
	//log = getLoggerInstance()
	//log.SetLogLevel(3)
	//log.Log("This is a log message")

	// create several goroutines that try to get the logger instance
	// concurrently
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			log := getLoggerInstance()
			log.SetLogLevel(i)
			log.Log("This is a log message")
		}()
	}
	_, _ = fmt.Scanln()
}
