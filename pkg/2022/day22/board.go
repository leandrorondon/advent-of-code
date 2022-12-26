package day22

const (
	Void  = 0
	Space = ' '
	Wall  = '#'
	Tile  = '.'

	Right = '>'
	Left  = '<'
	Down  = 'v'
	Up    = '^'
)

var rotation = map[byte]map[byte]byte{
	Right: {
		Clockwise:        Down,
		Counterclockwise: Up,
	},
	Down: {
		Clockwise:        Left,
		Counterclockwise: Right,
	},
	Left: {
		Clockwise:        Up,
		Counterclockwise: Down,
	},
	Up: {
		Clockwise:        Right,
		Counterclockwise: Left,
	},
}

type Position struct {
	Pos  Coordinate
	Face byte
}

func (p *Position) Rotate(dir byte) {
	p.Face = rotation[p.Face][dir]
}

func (p *Position) Move(pos Position) {
	p.Pos = pos.Pos
	p.Face = pos.Face
}

type Board struct {
	grid         [][]byte
	instructions []Instruction
	wrapper      Wrapper
}

type Wrapper interface {
	Password(pos Position) int
	Next(pos *Position) Position
}

func NewBoard(m [][]byte, instructions []Instruction, wrapper Wrapper) Board {
	return Board{
		grid:         m,
		instructions: instructions,
		wrapper:      wrapper,
	}
}

func (b Board) Run() int {
	col := 0
	for i, v := range b.grid[0] {
		if v == Tile {
			col = i
			break
		}
	}

	pos := Position{Coordinate{0, col}, Right}

	for _, instr := range b.instructions {
		b.execute(instr, &pos)
	}

	return b.wrapper.Password(pos)
}

func (b Board) execute(instruction Instruction, pos *Position) {
	if instruction.IsRotation() {
		pos.Rotate(instruction.Action)
		return
	}

	for i := 0; i < instruction.Steps; i++ {
		next := b.wrapper.Next(pos)
		nextTile := b.grid[next.Pos.Row][next.Pos.Col]
		if nextTile == Wall {
			break
		} else if nextTile == Tile {
			pos.Move(next)
		}
	}
}
