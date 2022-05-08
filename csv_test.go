package nnhelper

import (
	"testing"
)

// Reading existing csv file
func TestReadCsvFile(t *testing.T) {
	_, err := ReadCsvFile("examples/sam01/sam01_inp.csv")
	if err != nil {
		t.Errorf("Can't read csv file, error = %s; want nil", err.Error())
		return
	}
}

// Reading non-existing csv file
func TestReadCsvFile2(t *testing.T) {
	_, err := ReadCsvFile("examples/sam01/sam01_inp_notexists.csv")
	if err == nil {
		t.Errorf("Does not return error when read not existing csv file, "+
			"error = %v", err)
		return
	}
}
