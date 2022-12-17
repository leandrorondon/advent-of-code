package valve

import "fmt"

type Valve struct {
	Label   string
	Rate    int
	Tunnels []*Valve
}

func (v *Valve) Print() {
	fmt.Println(v.String())
}

func (v *Valve) String() string {
	s := fmt.Sprintf("{%s %d [", v.Label, v.Rate)
	for _, p := range v.Tunnels {
		s += fmt.Sprintf("%s|%d ", p.Label, p.Rate)
	}
	s = fmt.Sprintf("%s]}", s[:len(s)-1])

	return s
}
