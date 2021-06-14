package constants

import (
	"go-jvm/ch04/rtda"
	"go-jvm/ch05/instructions/base"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {
	//do nothing
}
