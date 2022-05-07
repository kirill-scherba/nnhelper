// Copyright 2021-2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example sam01 description:
//
// The task of this neural network is to decide what the game character should 
// do, based on 3 parameters:
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
// Examples for traning:
//
// 	Health	Weapons	The enemies	Decision
// 	50		1		1			Attack
// 	90		1		2			Attack
// 	80		0		1			Attack
// 	30		1		1			Steal
// 	60		1		2			Steal
// 	40		0		1			Steal
// 	90		1		7			Run away
// 	60		1		4			Run away
// 	10		0		1			Run away
// 	60		1		0			Nothing to do
// 	100		0		0			Nothing to do
//
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kirill-scherba/nnhelper"
)

const (
	SAM01_NN  = "sam01.nn"
	SAM01_INP = "sam01_inp.csv"
	SAM01_TAR = "sam01_tar.csv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	humanAnswers := []string{"Attack", "Steal", "Run away", "Nothing to do"}

	// Create NN if it file does not exists
	if _, err := os.Stat(SAM01_NN); errors.Is(err, os.ErrNotExist) {
		log.Println("Create", SAM01_NN, "neural network")
		nnhelper.Create(3, 16, 4, false, SAM01_INP, SAM01_TAR, SAM01_NN, true)
	}

	// Load neural network from file
	nn := nnhelper.Load(SAM01_NN)

	// Input values:
	// hp - heals (0.1 - 1.0)
	// weapon - weapon present (0 - no, 1 - yes)
	// enemyCount - enemy count
	var hp float64 = 0.7
	var weapon float64 = 1.0
	var enemyCount float64 = 1.0

	// Get answer from NN (weight array)
	out := nn.Answer(hp, weapon, enemyCount)
	answer := nn.AnswerToHuman(out, humanAnswers)

	// Print answer
	fmt.Println(hp, weapon, enemyCount, answer, out)
}
