package cpu

type memory struct {
	ram [4096]uint16
}

func (m memory) getInstruction(pc uint16) uint16 {
	return m.ram[pc];
}