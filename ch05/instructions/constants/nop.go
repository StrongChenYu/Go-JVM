package constants

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {
	//do nothing
}


