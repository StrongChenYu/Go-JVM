package math

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

type IMUL struct { base.NoOperandsInstruction }
type LMUL struct { base.NoOperandsInstruction }
type DMUL struct { base.NoOperandsInstruction }
type FMUL struct { base.NoOperandsInstruction }

func (self *IMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	result := op1 * op2
	stack.PushInt(result)
}

func (self *LMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopLong()
	op1 := stack.PopLong()

	result := op1 * op2
	stack.PushLong(result)
}

func (self *DMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopDouble()
	op1 := stack.PopDouble()

	result := op1 * op2
	stack.PushDouble(result)
}

func (self *FMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopFloat()
	op1 := stack.PopFloat()

	result := op1 * op2
	stack.PushFloat(result)
}

