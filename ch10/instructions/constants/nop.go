package constants

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
