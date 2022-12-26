package day22

const (
	Clockwise        = 'R'
	Counterclockwise = 'L'
)

type Instruction struct {
	Action byte
	Steps  int
}

func (ins Instruction) IsRotation() bool {
	return ins.Action == Clockwise || ins.Action == Counterclockwise
}
