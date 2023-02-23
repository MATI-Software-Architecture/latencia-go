package cmd

import (
	"fmt"
	"time"

	"github.com/mati/latencia/schema"
)

// Channels with buffer for tasks -> Queue / 10000 slots per task
var PriorityCh = make(chan schema.Task, 100000)
var ClassifyCh = make(chan schema.Task, 100000)
var NotifyCh = make(chan schema.Task, 100000)
var SendNotifyCh = make(chan schema.Task, 100000)
var SaveCh = make(chan schema.Task, 100000)

// Go routines -> lightweight threads / 1000 go routines per task
func StartTasks() {
	for i := 0; i < 10000; i++ {
		go PriorityTask(PriorityCh)
		go ClassifyTask(ClassifyCh)
		go NotifyTask(NotifyCh)
		go SendNotifyTask(SendNotifyCh)
		go SaveTask(SaveCh)
	}
}

// Validate the signal priority and send it to the next task
func PriorityTask(channel chan schema.Task) {
	for {
		signal := <-channel
		if signal.Body.Panic {
			signal.Emergency = []string{"panic"}
			NotifyCh <- signal
		} else {
			ClassifyCh <- signal
		}
	}
}

// check the signal for emergencies and warnings and send it to the next task
func ClassifyTask(channel chan schema.Task) {
	for {
		signal := <-channel
		emergencies, warnings := Classify(signal)
		switch {
		case emergencies != nil:
			signal.Emergency = emergencies
			signal.Warning = warnings
			NotifyCh <- signal
		case warnings != nil:
			signal.Warning = warnings
			NotifyCh <- signal
		default:
			SaveCh <- signal
		}
	}
}

// check the signal for endpoints and send it to the next tasks
func NotifyTask(channel chan schema.Task) {
	for {
		signal := <-channel
		signal.Endpoints = Endpoints(signal)
		SendNotifyCh <- signal
		SaveCh <- signal
	}
}

// send the signal to the endpoints
func SendNotifyTask(channel chan schema.Task) {
	for {
		signal := <-channel
		if signal.Elapsed > 100*time.Millisecond {
			fmt.Println("Slow task")
		}
		// for _, endpoint := range signal.Endpoints {
		// 	SendNotify(endpoint, signal)
		// }
	}
}

// save the signal in the database
func SaveTask(channel chan schema.Task) {
	for {
		signal := <-channel
		signal.Elapsed = time.Since(signal.Start)
		// Save(signal)
	}
}
