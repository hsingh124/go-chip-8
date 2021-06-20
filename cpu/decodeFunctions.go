package cpu

func decodeInstruction(instruction uint16) (uint16, [3]uint16) {
	return instruction >> 12, [3]uint16{instruction >> 8 & 15, instruction >> 4 & 15, instruction & 15}
}