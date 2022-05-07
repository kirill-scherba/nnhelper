package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kirill-scherba/nnhelper"
)

const (
	SAM02_NN  = "sam02.nn"
	SAM02_INP = "sam02_inp.csv"
	SAM02_TAR = "sam02_tar.csv"
)

// Time of day sample

var timeOfDay = flag.String("time", "", "time of day")
var timeNow = flag.Bool("now", false, "time of day now")

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	humanAnswers := []string{"Night", "Morning", "Day", "Evening"}

	// Create NN if it file does not exists
	if _, err := os.Stat(SAM02_NN); errors.Is(err, os.ErrNotExist) {
		log.Println("Create", SAM02_NN, "neural network")
		nnhelper.Create(1, 64, 4, false, SAM02_INP, SAM02_TAR, SAM02_NN, true)
	}

	// Load neural network from file
	nn := nnhelper.Load(SAM02_NN)

	// Check Input parameters:
	// timeArray (00:00 - 23:59)
	var timeArray []string
	switch {
	case len(*timeOfDay) > 0:
		timeArray = []string{*timeOfDay}
	case *timeNow:
		now := time.Now()
		timeArray = []string{fmt.Sprintf("%d:%02d", now.Hour(), now.Minute())}
	default:
		timeArray = []string{"10:11", "15:23", "18:20", "1:45"}
	}

	// Get answers from NN (weight array)
	for i := range timeArray {
		t, _ := nnhelper.TimeToFloat(timeArray[i])
		out := nn.Answer(t)
		answer := nn.AnswerToHuman(out, humanAnswers)

		// Print answer
		fmt.Printf("%s\t%s\t%v\n", timeArray[i], answer, out)
	}
}
