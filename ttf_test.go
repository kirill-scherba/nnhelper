package nnhelper

import (
	"fmt"
	"testing"
)

// Test time to float convert

func TestTimeToFloat(t *testing.T) {

	var times = []string{
		"16:33", "09:30", "09:30:30", "09:31", "22",
	}

	for i := range times {
		got := times[i]
		out, err := TimeToFloat(got)
		fmt.Printf("%-8s %9.6f\t%v\n", got, out, err)
		if err != nil {
			t.Errorf("Error = %s; want nil", err.Error())
			return
		}
	}
}
