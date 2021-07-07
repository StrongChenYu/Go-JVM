package stack

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
)

type POP struct{ base.NoOperandsInstruction }

//pop两个值出来
type POP2 struct{ base.NoOperandsInstruction }

func (p *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (p *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
