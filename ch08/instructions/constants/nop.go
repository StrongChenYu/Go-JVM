package constants

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
