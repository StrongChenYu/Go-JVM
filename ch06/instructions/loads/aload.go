package loads

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

//把index位置的localVars处的变量加载到操作数栈里面
type ALOAD struct { base.Index8Instruction }
type ALOAD_0 struct { base.NoOperandsInstruction }
type ALOAD_1 struct { base.NoOperandsInstruction }
type ALOAD_2 struct { base.NoOperandsInstruction }
type ALOAD_3 struct { base.NoOperandsInstruction }

func (self *ALOAD) Execute(frame *rtda.Frame)  {
	_aload(frame, self.Index)
}

func (self *ALOAD_0) Execute(frame *rtda.Frame)  {
	_aload(frame, 0)
}

func (self *ALOAD_1) Execute(frame *rtda.Frame)  {
	_aload(frame, 1)
}

func (self *ALOAD_2) Execute(frame *rtda.Frame)  {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *rtda.Frame)  {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint)  {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

