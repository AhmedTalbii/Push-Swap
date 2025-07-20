package solve

import (
	"fmt"

	"push-swap/instructions"
)

var instructionsArr []string

var StacksData = &instructions.Stacks{
	A: []int{2, 1, 3, 6, 8, 5},
	B: []int{},
}

func minIndex(a *[]int) int {
	minIndex := 0
	for i := 0; i < len(*a); i++ {
		if (*a)[minIndex] > (*a)[i] {
			minIndex = i
		}
	}
	return minIndex
}

func Solve(a, b *[]int, maxLen int) []string {
	for {
		if len(*a) == 0 {
			break
		}

		// search for the easiest way to make it in b
		if minIndex(a) == 0 {
			instructionsArr = append(instructionsArr, "pb")
			StacksData.Pb()
		} else if minIndex(a) == 1 {
			instructionsArr = append(instructionsArr, "sa")
			StacksData.Sa()
		} else if minIndex(a) > len(*a)/2 {
			for minIndex(a) > len(*a)/2 {
				instructionsArr = append(instructionsArr, "rra")
				StacksData.Rra()
			}
		} else if minIndex(a) <= len(*a)/2 {
			for minIndex(a) != 0 {
				instructionsArr = append(instructionsArr, "ra")
				StacksData.Ra()
			}
		}
		fmt.Println(StacksData)
	}

	for len(*b) > 0 {
		instructionsArr = append(instructionsArr, "pa")
		StacksData.Pa()
		fmt.Println(StacksData)
	}

	return instructionsArr
}
