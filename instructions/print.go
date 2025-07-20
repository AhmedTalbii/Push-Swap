package instructions

import (
	"fmt"
)

func (s *Stacks) PrintVertical() {
	lenA, lenB := len(s.A), len(s.B)

	if lenA > lenB {
		bStart := lenA - lenB
		counter := 0
		for counter < bStart {
			fmt.Println(s.A[counter])
			counter++
		}
		for i := bStart; i < lenA; i++ {
			fmt.Println(s.A[i], s.B[i-bStart])
		}
	} else if lenA < lenB {
		aStart := lenB - lenA
		counter := 0
		for counter < aStart {
			fmt.Println(" ", s.B[counter])
			counter++
		}
		for i := aStart; i < lenB; i++ {
			fmt.Println(s.A[i-aStart], s.B[i])
		}
	} else {
		for i := 0; i < lenB; i++ {
			fmt.Println(s.A[i], s.B[i])
		}
	}
	fmt.Println("= =")
	fmt.Println("a b")
}
