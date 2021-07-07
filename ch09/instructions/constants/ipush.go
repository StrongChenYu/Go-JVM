package constants

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
)

/**
The immediate byte is sign-extended to an int value.
That value is pushed onto the operand stack.
*/
type BIPUSH struct{ val int8 }

/**
  The immediate unsigned byte1 and byte2 values are assembled into an intermediate short,
  where the value of the short is (byte1 << 8) | byte2.
  The intermediate value is then sign-extended to an int value.
  That value is pushed onto the operand stack.
*/
type SIPUSH struct{ val int16 }

func (self *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	val := int32(self.val)
	frame.OperandStack().PushInt(val)
}

func (self *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	val := int32(self.val)
	frame.OperandStack().PushInt(val)
}
