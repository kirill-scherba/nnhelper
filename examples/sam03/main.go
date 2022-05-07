package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kirill-scherba/nnhelper"
)

const (
	SAM03_NN  = "sam03.nn"
	SAM03_INP = "sam03_inp.csv"
	SAM03_TAR = "sam03_tar.csv"
)

func main() {

	humanAnswers := []string{"Plus", "Minus"}

	// Create NN if it file does not exists
	if _, err := os.Stat(SAM03_NN); errors.Is(err, os.ErrNotExist) {
		log.Println("Create", SAM03_NN, "neural network")
		nnhelper.Create(2, 4, 2, false, SAM03_INP, SAM03_TAR, SAM03_NN, true)
	}

	// Load neural network from file
	nn := nnhelper.Load(SAM03_NN)

	const (
		PLUS  = 1.0
		MINUS = -1.0
	)

	// Intput array for testing
	in := [][]float64{
		{PLUS, PLUS},   // Plus * Plus = Plus
		{PLUS, MINUS},  // Plus * Minus = Minus
		{MINUS, PLUS},  // Minus * Plus = Minus
		{MINUS, MINUS}, // Minus * Minus = Plus
		{3000, -0.001}, // Minus * Plus = Minus
	}
	for i := range in {
		out := nn.Answer(in[i]...)
		answer, _ := nn.AnswerToHuman(out, humanAnswers)
		fmt.Println(in[i], answer, out)
	}
}
