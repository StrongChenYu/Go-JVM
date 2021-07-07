package control

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
)

type RETURN struct{ base.NoOperandsInstruction }

type ARETURN struct{ base.NoOperandsInstruction }
type LRETURN struct{ base.NoOperandsInstruction }
type FRETURN struct{ base.NoOperandsInstruction }
type DRETURN struct{ base.NoOperandsInstruction }
type IRETURN struct{ base.NoOperandsInstruction }

//void return
func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(val)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopLong()
	invokeFrame.OperandStack().PushLong(val)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokeFrame.OperandStack().PushFloat(val)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokeFrame.OperandStack().PushDouble(val)
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.CurrentFrame()
	val := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(val)
}
