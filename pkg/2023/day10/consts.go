package day10

type Direction string

const (
	up    Direction = "u"
	down  Direction = "d"
	left  Direction = "l"
	right Direction = "r"
)

const (
	vertical    = "|"
	horizontal  = "-"
	topleft     = "F"
	topright    = "7"
	bottomleft  = "L"
	bottomright = "J"
	start       = "S"
	ground      = "."
)

var directions = []Direction{up, down, left, right}

var directionsToCheck = map[string][]Direction{
	vertical:    {up, down},
	horizontal:  {left, right},
	topleft:     {down, right},
	topright:    {down, left},
	bottomleft:  {up, right},
	bottomright: {up, left},
	start:       {up, down, left, right},
}

var oppositeConnection = map[Direction][]string{
	up:    {vertical, topleft, topright, start},
	down:  {vertical, bottomleft, bottomright, start},
	left:  {horizontal, topleft, bottomleft, start},
	right: {horizontal, topright, bottomright, start},
}
