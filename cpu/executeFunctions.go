package cpu

func (c Cpu) clearDisplay() {

}

func (c Cpu) returnFromASubroutine() {

}

func (c Cpu) jump(nnn uint16) {
	c.regs.programCounter = nnn
}

func (c Cpu) callSubRoutine(nnn uint16) {
	c.stack[c.regs.stackPointer] = c.regs.programCounter
	c.regs.stackPointer++
	c.regs.programCounter = nnn
}

func (c Cpu) skipNextInstructionIfEqualImmediate(x uint16, kk uint16) {
	if c.regs.generalRegs[x] == kk {
		c.regs.programCounter += 2
	}
}

func (c Cpu) skipNextInstructionIfNotEqualImmediate(x uint16, kk uint16) {
	if c.regs.generalRegs[x] != kk {
		c.regs.programCounter += 2
	}
}

func (c Cpu) skipNextInstructionIfEqualReg(x uint16, y uint16) {
	if c.regs.generalRegs[x] == c.regs.generalRegs[y] {
		c.regs.programCounter += 2
	}
}

func (c Cpu) loadImmediate(x uint16, kk uint16) {
	c.regs.generalRegs[x] = kk
}

func (c Cpu) addImmediate() {

}

func (c Cpu) loadReg() {

}

func (c Cpu) orReg() {
	
}

func (c Cpu) andReg() {
	
}

func (c Cpu) xorReg() {
	
}

func (c Cpu) addReg() {
	
}

func (c Cpu) subReg() {
	
}

func (c Cpu) shiftRight() {

}

func (c Cpu) subN() {

}

func (c Cpu) shiftLeft() {

}

func (c Cpu) skipNextInstructionIfNotEqualReg() {

}