package math

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

//异或运算
type IXOR struct { base.NoOperandsInstruction }
type LXOR struct { base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op1 := stack.PopInt()
	op2 := stack.PopInt()

	result := op1 ^ op2
	stack.PushInt(result)
}

func (self *LXOR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op1 := stack.PopLong()
	op2 := stack.PopLong()

	result := op1 ^ op2
	stack.PushLong(result)
}