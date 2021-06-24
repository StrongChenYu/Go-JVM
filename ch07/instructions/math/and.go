package math

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

//and布尔运算
type IAND struct { base.NoOperandsInstruction }
type LAND struct { base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op1 := stack.PopInt()
	op2 := stack.PopInt()

	result := op1 & op2
	stack.PushInt(result)
}

func (self *LAND) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op1 := stack.PopLong()
	op2 := stack.PopLong()

	result := op1 & op2
	stack.PushLong(result)
}