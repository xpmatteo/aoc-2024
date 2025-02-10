package day17

import (
	"log"
	"strconv"
	"strings"
)

type Machine struct {
	A, B, C int
	Program []int
}

func (m *Machine) Execute() string {
	var output []string
	pc := 0
	for pc < len(m.Program) {
		opcode := m.Program[pc]
		operand := m.Program[pc+1]
		switch opcode {
		case 0:
			m.A = m.A / (1 << m.combo(operand))
		case 1:
			m.B = m.B ^ operand
		case 2:
			m.B = m.combo(operand) % 8
		case 4:
			m.B = m.B ^ m.C
		case 5:
			output = append(output, strconv.Itoa(m.combo(operand)%8))
		default:
			log.Fatal("Unknown opcode: ", opcode)
		}
		pc += 2
	}
	return strings.Join(output, ",")
}

func (m *Machine) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return m.A
	case 5:
		return m.B
	case 6:
		return m.C
	default:
		log.Fatal("Unknown operand ", operand)
		return -1
	}
}
