// Copyright 2021 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Neural Network Helper package.
// This module Read CSV files and convert its values from string to float.

package nnhelper

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

// ReadCsvFile read csv file and return float array
func ReadCsv(filePath string) (out [][]float64, err error) {
	csv, err := ReadCsvFile(filePath)
	if err != nil {
		return
	}

	out, err = CsvToFloat(csv)
	return
}

// ReadCsvFile read csv file and return string array
func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		e := "Unable to read input file " + filePath + ": " + err.Error()
		return nil, errors.New(e)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		e := "Unable to parse file as CSV for " + filePath + ": " + err.Error()
		return nil, errors.New(e)
	}

	return records, nil
}

// CsvToFloat convert string csv array to float
func CsvToFloat(in [][]string) (out [][]float64, err error) {
	for _, fa := range in {
		var va []float64
		for _, f := range fa {
			var val float64
			if strings.Contains(f, ":") {
				val, err = TimeToFloat(f)
			} else {
				val, err = strconv.ParseFloat(strings.TrimSpace(f), 64)
			}
			if err != nil {
				return
			}
			va = append(va, val)
		}
		out = append(out, va)
	}
	return
}

// TimeToFloat convert string time of day to float value
func TimeToFloat(time string) (out float64, err error) {
	s := strings.Split(time, ":")

	out, err = strconv.ParseFloat(s[0], 64)
	if err != nil {
		return 0, err
	}

	if len(s) > 1 {
		fract, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			return 0, err
		}
		out += fract / 60
	}

	if len(s) > 2 {
		fract, err := strconv.ParseFloat(s[2], 64)
		if err != nil {
			return 0, err
		}
		out += fract / (60 * 60)
	}

	return
}
