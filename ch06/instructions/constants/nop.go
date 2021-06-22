package constants

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {
	//do nothing
}


