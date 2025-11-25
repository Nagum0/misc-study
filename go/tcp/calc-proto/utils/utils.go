package utils

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Operation int
const (
	PLUS = iota
	MINUS
	MULT
	DIV
)

func (op Operation) String() string {
	switch op {
		case PLUS:
			return "PLUS"
		case MINUS:
			return "MINUS"
		case MULT:
			return "MULT"
		case DIV:
			return "DIV"
	}

	panic("Unknown operation...")
}

type CalcData struct {
	A int `yaml:"a"`
	B int `yaml:"b"`
	Op Operation `yaml:"op"`
}

func (c CalcData) Calc() int {
	switch c.Op {
		case PLUS:
			return c.A + c.B
		case MINUS:
			return c.A - c.B
		case MULT:
			return c.A * c.B
		case DIV:
			return c.A / c.B
	}
	
	panic("Unknown operation...")
}

func (c CalcData) ToYaml() []byte {
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	return yamlBytes
}

func (c *CalcData) FromYaml(yamlBytes []byte) {
	if err := yaml.Unmarshal(yamlBytes, c); err != nil {
		panic(err)
	}
}

func (c CalcData) String() string {
	return fmt.Sprintf("%v %v %v", c.A, c.Op, c.B)
}
