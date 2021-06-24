package math

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

//or 布尔运算
type IOR struct { base.NoOperandsInstruction }
type LOR struct { base.NoOperandsInstruction }


func (self *IOR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op1 := stack.PopInt()
	op2 := stack.PopInt()

	result := op1 | op2
	stack.PushInt(result)
}

func (self *LOR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op1 := stack.PopLong()
	op2 := stack.PopLong()

	result := op1 | op2
	stack.PushLong(result)
}