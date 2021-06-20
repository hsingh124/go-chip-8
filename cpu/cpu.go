package cpu

type Cpu struct {
	ram memory
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
	case 7:
	case 8:
	case 9:
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
	case 15:
	}
}
