// Copyright 2021 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Neural Network Helper package.
// Create Neural Network module.
package nnhelper

import (
	"fmt"

	"github.com/fxsjy/gonn/gonn"
)

// CreateNN create Neural Network
func CreateNN(inputCount, hiddenCount, outputCount int, regression bool,
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
