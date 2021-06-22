package loads

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

//把index位置的localVars处的变量加载到操作数栈里面
/*
The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
The local variable at index must contain an int. The value of the local variable at index is pushed onto the operand stack.
 */
type ILOAD struct { base.Index8Instruction }
type ILOAD_0 struct { base.NoOperandsInstruction }
type ILOAD_1 struct { base.NoOperandsInstruction }
type ILOAD_2 struct { base.NoOperandsInstruction }
type ILOAD_3 struct { base.NoOperandsInstruction }

func (self *ILOAD) Execute(frame *rtda.Frame)  {
	_iload(frame, self.Index)
}

func (self *ILOAD_0) Execute(frame *rtda.Frame)  {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame)  {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame)  {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame)  {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint)  {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
