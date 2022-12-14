package decoder

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Value struct {
	Value  *int
	Packet Packet
}

func (v *Value) UnmarshalJSON(data []byte) error {
	if data[0] == '[' {
		err := json.Unmarshal(data, &v.Packet)
		if err != nil {
			fmt.Println("err Packet:", string(data), err)
		}
		return err
	}

	err := json.Unmarshal(data, &v.Value)
	if err != nil {
		fmt.Println("err Value:", string(data), err)
	}
	return err
}

func (v *Value) Compare(v2 Value) int {
	if v.Value != nil && v2.Value != nil {
		if *v.Value < *v2.Value {
			return -1
		}
		if *v.Value > *v2.Value {
			return 1
		}
		return 0
	}

	if v.Value != nil && v2.Packet != nil {
		newV := &Value{Packet: []Value{*v}}
		return newV.Compare(v2)
	}

	if v.Packet != nil && v2.Value != nil {
		newV := Value{Packet: []Value{v2}}
		return v.Compare(newV)
	}

	return v.Packet.Compare(v2.Packet)
}

func (v Value) String() string {
	if v.Value != nil {
		return strconv.Itoa(*v.Value)
	}

	return v.Packet.String()
}
