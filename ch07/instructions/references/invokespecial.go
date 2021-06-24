package references

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

type INVOKE_SPECIAL struct { base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PopRef()
}
