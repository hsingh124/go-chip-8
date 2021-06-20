package cpu

type Cpu struct {
	ram memory
}

func decodeInstruction(instruction uint16) (uint16, [3]uint16) {
	return instruction >> 12, [3]uint16{instruction >> 8 & 15, instruction >> 4 & 15, instruction & 15}
}

func (c Cpu) executeInstruction(instruction uint16) {
	opcode, vars := decodeInstruction(instruction)

	switch opcode {
	case 0:
		// ignoring 0nnn for now
		if vars[2] == 0 {
			c.clearDisplay()
		} else {
			c.returnFromASubroutine()
		}
	case 1:
		c.jump()
	case 2:
		c.callSubRoutine()
	case 3:
		c.skipNextInstructionIfEqualImmediate()
	case 4:
		c.skipNextInstructionIfNotEqualImmediate()
	case 5:
		c.skipNextInstructionIfEqualReg()
	case 6:
		c.loadImmediate()
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
		case 7:
		case 14:
		default:
		}
	case 9:
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
	case 15:
	default:
	}
}
