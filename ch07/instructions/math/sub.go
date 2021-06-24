package math

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

type ISUB struct { base.NoOperandsInstruction }
type LSUB struct { base.NoOperandsInstruction }
type DSUB struct { base.NoOperandsInstruction }
type FSUB struct { base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	result := op1 - op2
	stack.PushInt(result)
}

func (self *LSUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopLong()
	op1 := stack.PopLong()

	result := op1 - op2
	stack.PushLong(result)
}

func (self *DSUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopDouble()
	op1 := stack.PopDouble()

	result := op1 - op2
	stack.PushDouble(result)
}

func (self *FSUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	op2 := stack.PopFloat()
	op1 := stack.PopFloat()

	result := op1 - op2
	stack.PushFloat(result)
}
