package cpu

type registers struct {
	programCounter uint16
	stackPointer uint8
	generalRegs [16]uint16
}