package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fxsjy/gonn/gonn"
	"github.com/kirill-scherba/nnhelper"
)

const (
	SAM01_NN  = "sam01.nn"
	SAM01_INP = "sam01_inp.csv"
	SAM01_TAR = "sam01_tar.csv"
)

// Example description:
//
// The task of the neural network is to decide what the character should do,
// based on 3 parameters:
//
// 	- Amount of health (from 1 to 100)
// 	- The presence of weapons
// 	- Number of enemies
//
// Depending on the outcome, one of the following decisions may be taken:
//
// 	- Attack
// 	- Steal
// 	- Run away
// 	- Nothing to do
//
// Examples:
//
// 	Health	Weapons	The enemies	Decision
// 	50		one		one			Attack
// 	90		one		2			Attack
// 	80		0		one			Attack
// 	thirty	one		one			Steal
// 	60		one		2			Steal
// 	40		0		one			Steal
// 	90		one		7			Run away
// 	60		one		four		Run away
// 	ten		0		one			Run away
// 	60		one		0			Nothing to do
// 	100		0		0			Nothing to do
//

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Create NN if it file does not exists
	if _, err := os.Stat(SAM01_NN); errors.Is(err, os.ErrNotExist) {
		log.Println("Create", SAM01_NN, "neural network")
		nnhelper.CreateNN(3, 16, 4, false, SAM01_INP, SAM01_TAR, SAM01_NN, true)
	}

	// Load neural network from file
	// log.Println("Load", SAM01_NN, "neural network")
	nn := gonn.LoadNN(SAM01_NN)

	// Input values:
	// hp - heals (0.1 - 1.0)
	// weapon - weapon present (0 - no, 1 - yes)
	// enemyCount - enemy count
	var hp float64 = 0.7
	var weapon float64 = 1.0
	var enemyCount float64 = 1.0

	// Get ansver from NN (weight array)
	// log.Println("Execute", SAM01_NN, "neural network")
	out := nn.Forward([]float64{hp, weapon, enemyCount})

	// Print answer
	fmt.Println(hp, weapon, enemyCount, GetResult(out), out)

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
		return "Attack"
	case 1:
		return "Steal"
	case 2:
		return "Run away"
	case 3:
		return "Nothing to do"
	}
	return ""
}
