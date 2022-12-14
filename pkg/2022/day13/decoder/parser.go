package decoder

import (
	"encoding/json"
)

type Parser struct {
	scanner Scanner
}

type Scanner interface {
	Scan() bool
	Text() string
}

func NewParser(s Scanner) *Parser {
	return &Parser{
		scanner: s,
	}
}

func (p *Parser) ParseSignal() Signal {
	var signal Signal

	for p.scanner.Scan() {
		text := p.scanner.Text()
		if text == "" {
			continue
		}

		packet := p.ParsePacket(text)
		signal = append(signal, packet)
	}

	return signal
}

func (p *Parser) ParsePacket(text string) Packet {
	var packet Packet
	_ = json.Unmarshal([]byte(text), &packet)
	return packet
}
