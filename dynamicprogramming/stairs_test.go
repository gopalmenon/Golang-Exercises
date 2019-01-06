package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestSteps(t *testing.T) {

	fmt.Printf("%d ways to climb stairs with %d steps", HowManyStepWays(10), 10)

	if ways := HowManyStepWays(1); ways != 1 {
		t.Errorf("Expected 1 way to climb stairs with 1 step, but got %d", ways)
	}

	if ways := HowManyStepWays(2); ways != 2 {
		t.Errorf("Expected 3 ways to climb stairs with 2 steps, but got %d", ways)
	}

	if ways := HowManyStepWays(3); ways != 4 {
		t.Errorf("Expected 4 ways to climb stairs with 3 steps, but got %d", ways)
	}

}
