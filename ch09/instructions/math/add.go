package math

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
)

//加法指令
type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }
type DADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	result := op1 + op2
	stack.PushInt(result)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopLong()
	op1 := stack.PopLong()

	result := op1 + op2
	stack.PushLong(result)
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopDouble()
	op1 := stack.PopDouble()

	result := op1 + op2
	stack.PushDouble(result)
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopFloat()
	op1 := stack.PopFloat()

	result := op1 + op2
	stack.PushFloat(result)
}
