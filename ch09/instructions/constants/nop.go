package constants

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
