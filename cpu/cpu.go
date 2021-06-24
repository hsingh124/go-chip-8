package cpu

type Cpu struct {
	ram [4096]uint16
	stack [16]uint16
	regs registers
}

func decodeInstruction(instruction uint16) (uint16, [3]uint16) {
	return instruction >> 12, [3]uint16{instruction >> 8 & 15, instruction >> 4 & 15, instruction & 15}
}

func getNNN(instruction uint16) uint16 {
	// bitwise AND of instruction with 0000 1111 1111 1111 which is 4095
	return instruction & 4095
}

func getKK(instruction uint16) uint16 {
	// bitwise AND of instruction with 0000 0000 1111 1111 which is 255
	return instruction & 255
}

func (c Cpu) executeInstruction(instruction uint16) {
	opcode, vars := decodeInstruction(instruction)
	nnn := getNNN(instruction)
	kk := getKK(instruction)

	switch opcode {
	case 0:
		// ignoring 0nnn for now
		if vars[2] == 0 {
			c.clearDisplay()
		} else {
			c.returnFromASubroutine()
		}
	case 1:
		c.jump(nnn)
	case 2:
		c.callSubRoutine(nnn)
	case 3:
		c.skipNextInstructionIfEqualImmediate(vars[0], kk)
	case 4:
		c.skipNextInstructionIfNotEqualImmediate(vars[0], kk)
	case 5:
		c.skipNextInstructionIfEqualReg(vars[0], vars[1])
	case 6:
		c.loadImmediate(vars[0], kk)
	case 7:
		c.addImmediate()
	case 8:
		switch vars[2] {
		case 0:
			c.loadReg()
		case 1:
			c.orReg()
		case 2:
			c.andReg()
		case 3:
			c.xorReg()
		case 4:
			c.addReg()
		case 5:
			c.subReg()
		case 6:
			c.shiftRight()
		case 7:
			c.subN()
		case 14:
			c.shiftRight()
		default:
		}
	case 9:
		c.skipNextInstructionIfNotEqualReg()
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
	case 15:
	default:
	}
}
