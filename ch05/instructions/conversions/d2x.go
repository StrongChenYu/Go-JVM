package conversions

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

//double 型数据转换
type D2L struct { base.NoOperandsInstruction }
type D2F struct { base.NoOperandsInstruction }
type D2I struct { base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushLong(int64(val))
}

func (self *D2F) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushFloat(float32(val))
}

func (self *D2I) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushInt(int32(val))
}