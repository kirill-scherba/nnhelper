package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fxsjy/gonn/gonn"
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

	// Create NN if it file does not exists
	if _, err := os.Stat(SAM02_NN); errors.Is(err, os.ErrNotExist) {
		log.Println("Create", SAM02_NN, "neural network")
		nnhelper.CreateNN(1, 64, 4, false, SAM02_INP, SAM02_TAR, SAM02_NN, true)
	}

	// Load neural network from file
	// log.Println("Load", SAM02_NN, "neural network")
	nn := gonn.LoadNN(SAM02_NN)

	// Input values:
	// timeArray (00:00 - 23:59)
	var timeArray []string
	if *timeOfDay != "" {
		timeArray = []string{*timeOfDay}
	} else if *timeNow {
		now := time.Now()
		timeArray = []string{fmt.Sprintf("%d:%02d", now.Hour(), now.Minute())}
	} else {
		timeArray = []string{"10:11", "15:23", "18:20", "1:45"}
	}

	// Get ansvers from NN (weight array)
	// log.Println("Execute", SAM02_NN, "neural network")
	for i := range timeArray {
		t, _ := nnhelper.TimeToFloat(timeArray[i])
		out := nn.Forward([]float64{t})

		// Print answer
		fmt.Printf("%s\t%s\t%v\n", timeArray[i], GetResult(out), out)
	}

	// log.Println("All done")
}

func GetResult(output []float64) string {
	max := -99999.0
	pos := -1
	// Get max weight position
	for i, value := range output {
		if value > max {
			max = value
			pos = i
		}
	}

	// Get result depend on position
	switch pos {
	case 0:
		return "Night"
	case 1:
		return "Morning"
	case 2:
		return "Day"
	case 3:
		return "Evening"
	}
	return ""
}
