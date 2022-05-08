// Copyright 2021-2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Neural Network Helper package.
// Create and Load Neural Network from file or string. Read CSV files and
// convert its values from string to float
package nnhelper

import (
	"encoding/json"
	"fmt"

	"github.com/fxsjy/gonn/gonn"
)

// Create Neural Network
func Create(inputCount, hiddenCount, outputCount int, regression bool,
	inpupCsv, targetCsv, resultNN string, print ...bool) {

	nn := gonn.DefaultNetwork(inputCount, hiddenCount, outputCount, regression)

	input, _ := ReadCsv(inpupCsv)
	if len(print) > 0 && print[0] {
		fmt.Println("Input: ", input)
	}

	target, _ := ReadCsv(targetCsv)
	if len(print) > 0 && print[0] {
		fmt.Println("Target:", target)
	}

	nn.Train(input, target, 100000)

	gonn.DumpNN(resultNN, nn)
}

// MaxFloatPosition return position of maximum weight value in float array,
// or -1 if array is empty
func MaxFloatPosition(in []float64) int {
	var max float64
	pos := -1
	// Get max weight position
	for i, value := range in {
		if i == 0 || value > max {
			max = value
			pos = i
		}
	}
	return pos
}

// Neural Network
type NeuralNetwork struct {
	*gonn.NeuralNetwork
}

// Load neural network from file
func Load(fileName string) *NeuralNetwork {
	nn := gonn.LoadNN(fileName)
	return &NeuralNetwork{nn}
}

// Load neural network from string
func LoadFromString(nnstrig string) *NeuralNetwork {
	nn := &gonn.NeuralNetwork{}
	err := json.Unmarshal([]byte(nnstrig), nn)
	if err != nil {
		panic(err)
	}
	return &NeuralNetwork{nn}
}

// Get answer from neural network (get weight array), return output array
func (nn *NeuralNetwork) Answer(in ...float64) (out []float64) {
	return nn.Forward(in)
}

// AnswerToHuman translate nn answer to human answer, return human string value
// and index in output array
func (nn *NeuralNetwork) AnswerToHuman(out []float64, human []string) (string, int) {
	pos := MaxFloatPosition(out)
	if pos >= 0 && pos < len(human) {
		return human[pos], pos
	}
	return "", pos
}
