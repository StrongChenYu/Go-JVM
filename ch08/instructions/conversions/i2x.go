package conversions

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

//int 型数据转换
type I2L struct{ base.NoOperandsInstruction }
type I2F struct{ base.NoOperandsInstruction }
type I2D struct{ base.NoOperandsInstruction }

//int -> byte
//The value on the top of the operand stack must be of type int.
//It is popped from the operand stack, truncated to a byte,
//then sign-extended to an int result. That result is pushed onto the operand stack.
//注意是sign-extended
type I2B struct{ base.NoOperandsInstruction }

//int to char，zero-extend
type I2C struct{ base.NoOperandsInstruction }
type I2S struct{ base.NoOperandsInstruction }

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(int8(val)))
}

func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	//注意这里的uint16,为什么转化成字符串是uint16
	stack.PushInt(int32(uint16(val)))
}

func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(int16(val)))
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushLong(int64(val))
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushFloat(float32(val))
}

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushDouble(float64(val))
}
