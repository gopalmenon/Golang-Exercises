//The number of ways a child can run up a staircase with n steps, where he can hop either 1 step, 2 steps or 3 steps at a time

package dynamicprogramming

func HowManyStepWays(steps int) int {

	waysCounts := make([]int, steps+1)
	return countWays(waysCounts, steps)

}

func countWays(waysCounts []int, steps int) int {

	if steps == 0 {
		return 1
	}

	if waysCounts[steps] != 0 {
		return waysCounts[steps]
	}

	var stepWays1, stepWays2, stepWays3 int
	if steps >= 3 {
		stepWays3 = countWays(waysCounts, steps-3)
	}
	if steps >= 2 {
		stepWays2 = countWays(waysCounts, steps-2)
	}
	if steps >= 1 {
		stepWays1 = countWays(waysCounts, steps-1)
	}

	waysCounts[steps] = stepWays1 + stepWays2 + stepWays3

	return waysCounts[steps]

}
