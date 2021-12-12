package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		fmt.Println(err)
		t.Errorf("Failed to run!")
	}
}
