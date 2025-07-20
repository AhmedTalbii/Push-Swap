package instructions

type Stacks struct {
	A []int
	B []int
}

// "Push the top element of stack B to stack A".
func (s *Stacks) Pa() {
	if len(s.B) < 1 {
		return
	}
	temp := []int{}
	temp = append(temp, s.B[0])
	temp = append(temp, s.A...)
	s.A = temp
	s.B = s.B[1:]
}

// "Push the top element of stack A to stack B".
func (s *Stacks) Pb() {
	if len(s.A) < 1 {
		return
	}
	temp := []int{}
	temp = append(temp, s.A[0])
	temp = append(temp, s.B...)
	s.B = temp
	s.A = s.A[1:]
}

// "Swap the first two elements of stack A".
func (s *Stacks) Sa() {
	if len(s.A) < 2 {
		return
	}
	s.A[0], s.A[1] = s.A[1], s.A[0]
}

// "Swap the first two elements of stack B".
func (s *Stacks) Sb() {
	if len(s.B) < 2 {
		return
	}
	s.B[0], s.B[1] = s.B[1], s.B[0]
}

// "Execute sa and sb".
func (s *Stacks) Ss() {
	s.Sa()
	s.Sb()
}

// "Rotate stack A (shift up all elements, first becomes last)".
func (s *Stacks) Ra() {
	temp := []int{}
	temp = append(temp, s.A[1:]...)
	temp = append(temp, s.A[0])
	s.A = temp
}

// "Rotate stack B (shift up all elements, first becomes last)".
func (s *Stacks) Rb() {
	temp := []int{}
	temp = append(temp, s.B[1:]...)
	temp = append(temp, s.B[0])
	s.B = temp
}

// "Execute ra and rb")
func (s *Stacks) Rr() {
	s.Ra()
	s.Rb()
}

// "Reverse rotate stack A (shift down all elements, last becomes first)".
func (s *Stacks) Rra() {
	temp := []int{}
	temp = append(temp, s.A[len(s.A)-1])
	temp = append(temp, s.A[:len(s.A)-1]...)
	s.A = temp
}

// "Reverse rotate stack B (shift down all elements, last becomes first)".
func (s *Stacks) Rrb() {
	temp := []int{}
	temp = append(temp, s.B[len(s.B)-1])
	temp = append(temp, s.B[:len(s.B)-1]...)
	s.B = temp
}

// "Execute rra and rrb".
func (s *Stacks) Rrr() {
	s.Rra()
	s.Rrb()
}
