package day15

import (
	"slices"
)

func HashP1(b []byte) int {
	sum := 0
	h := 0

	for i := range b {
		if b[i] == ',' {
			sum += h
			h = 0
			continue
		}

		h += int(b[i])
		h *= 17
		h %= 256
	}

	return sum + h
}

type lens struct {
	label string
	focal byte
}

func TotalP2(b []byte) int {
	var boxes [256][]lens
	h := 0
	var focal, op byte
	label := ""

	for i := range b {
		if b[i] == ',' {
			operation(&boxes, h, op, focal, label)
			h = 0
			label = ""
			focal = 0
			continue
		}

		if b[i] == '-' || b[i] == '=' {
			op = b[i]
			continue
		}

		if b[i] >= '1' && b[i] <= '9' {
			focal = b[i] - '0'
			continue
		}

		h += int(b[i])
		h *= 17
		h %= 256
		label += string(b[i])
	}
	operation(&boxes, h, op, focal, label)

	sum := 0
	for box := range boxes {
		for slot := range boxes[box] {
			sum += (box + 1) * (slot + 1) * int(boxes[box][slot].focal)
		}
	}

	return sum
}

func operation(boxes *[256][]lens, box int, op, focal byte, label string) {
	if op == '-' {
		remove(boxes, box, label)
		return
	}

	add(boxes, box, focal, label)
}

func add(boxes *[256][]lens, box int, focal byte, label string) {
	i := find(boxes[box], label)

	if i >= 0 {
		boxes[box][i].focal = focal
	} else {
		boxes[box] = append(boxes[box], lens{label, focal})
	}
}

func remove(boxes *[256][]lens, box int, label string) {
	if i := find(boxes[box], label); i >= 0 {
		boxes[box] = slices.Delete(boxes[box], i, i+1)
	}
}

func find(arr []lens, label string) int {
	for i := range arr {
		if arr[i].label == label {
			return i
		}
	}

	return -1
}
