package main

import (
	//"fmt"
	"fmt"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v: Month: %v Year: %v",
		label, t.Day(), t.Month(), t.Year())
}

func writeToChannel(channel chan<- string) {

	timer := time.NewTimer(time.Minute * 10)
	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second)
	}()
	Printfln("Waiting for initial duration...")

	// _ = <-time.After(time.Second * 2)
	<-timer.C

	Printfln("Initial duration elapsed")

	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		// time.Sleep(time.Second * 3)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	// отложенный запуск
	// time.AfterFunc(time.Second*5, func() { writeToChannel(nameChannel) })

	go writeToChannel(nameChannel)

	channelOpen := true
	for channelOpen {
		Printfln("Starting channel read")
		select {
		case name, ok := <-nameChannel:
			if !ok {
				channelOpen = false
				break
			} else {
				Printfln("Read name: %v", name)
			}
		case <-time.After(time.Second * 2):
			Printfln("Timeout")
		}
	}
}
