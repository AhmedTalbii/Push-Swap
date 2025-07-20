package main

import (
	"fmt"

	"push-swap/instructions"
)

var StacksData = &instructions.Stacks{
	A: []int{2, 1, 3, 6, 8, 5},
	B: []int{},
}

func main() {
	var name string
	for {
		instructions.ClearTerminal()
		StacksData.PrintVertical()
		fmt.Print("Enter your instruction : ")
		fmt.Scan(&name)
		switch name {
		case "pa":
			StacksData.Pa()
		case "pb":
			StacksData.Pb()
		case "sa":
			StacksData.Sa()
		case "sb":
			StacksData.Sb()
		case "ss":
			StacksData.Ss()
		case "ra":
			StacksData.Ra()
		case "rb":
			StacksData.Rb()
		case "rr":
			StacksData.Rr()
		case "rra":
			StacksData.Rra()
		case "rrb":
			StacksData.Rrb()
		case "rrr":
			StacksData.Rrr()
		case "exit":
			return
		default:
			fmt.Println("Unknown instruction:", name)
		}
	}
}
