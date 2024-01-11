// Copyright 2024 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example sam04 description:
//
// # Multiplication table
//
// Input parameters:
//
// - first digit (1-9) four float pins for first digit
// - second digit (1-9) four float pins for second digit
//
// Output parameters:
//
// - eight float pins for multiplication 9*9 = 0101 0001 B
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kirill-scherba/nnhelper"
)

const (
	SAM04_NN  = "sam04.nn"
	SAM04_INP = "sam04_inp.csv"
	SAM04_TAR = "sam04_tar.csv"
)

func main() {

	// Create input and target slices
	var input [][]float64
	var target [][]float64
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			input = append(input, append(makeBitArray(i), makeBitArray(j)...))
			target = append(target, makeBitArray(i*j, 8))
		}
	}

	// Repeat input and target slices for values with errors
	for range make([]int, 10) {
		input = append(input, append(makeBitArray(3), makeBitArray(3)...))
		target = append(target, makeBitArray(3*3, 8))

		input = append(input, append(makeBitArray(3), makeBitArray(6)...))
		target = append(target, makeBitArray(3*6, 8))

		input = append(input, append(makeBitArray(6), makeBitArray(6)...))
		target = append(target, makeBitArray(6*6, 8))
	}

	// Create NN if it file does not exists
	if _, err := os.Stat(SAM04_NN); errors.Is(err, os.ErrNotExist) {
		log.Println("Create", SAM04_NN, "neural network")
		nnhelper.CreateFromSlice(8, 24, 8, false, input, target, SAM04_NN, true)
	}

	// Load neural network from file
	nn := nnhelper.Load(SAM04_NN)

	// Check results
	var in [][]float64
	// Create input array
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			in = append(in, append(makeBitArray(i), makeBitArray(j)...))
		}
	}
	// Get and print results
	for i := range in {
		t := time.Now()
		out := nn.Answer(in[i]...)
		ms := float64(time.Since(t).Microseconds())/1000.00
		first, second, result := makeDigit(in[i][:4]), makeDigit(in[i][4:]),
			makeDigit(out)
		fmt.Printf("%v %d * %d = %d %v (%.3fms)\n",
			in[i],
			first,
			second,
			result,
			colorBool(first*second == result),
			ms,
		)
	}
}

// makeBitArray make bit array from int
func makeBitArray(num int, lens ...int) []float64 {
	l := 4
	if len(lens) > 0 {
		l = lens[0]
	}
	bitArray := make([]float64, l)
	for k := range bitArray {
		bitValue := (num >> k) & 1
		bitArray[l-k-1] = float64(bitValue)
	}
	return bitArray
}

// makeDigit make int from bit array
func makeDigit(bits []float64) int {
	digit := 0
	j := 0
	for i := len(bits) - 1; i >= 0; i-- {
		var b int
		if bits[i] > 0.5 {
			b = 1
		}
		digit |= b << j
		j++
	}
	return digit
}

// colorBool return color string for bool value
func colorBool(b bool) string {
	if b {
		return "\x1b[32mtrue\x1b[0m"
	}
	return "\x1b[31mfalse\x1b[0m"
}
