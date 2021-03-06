package stores

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

//将操作数栈顶Long类型数据放到位置为i的LocalVars
type DSTORE struct { base.Index8Instruction }
type DSTORE_0 struct { base.NoOperandsInstruction }
type DSTORE_1 struct { base.NoOperandsInstruction }
type DSTORE_2 struct { base.NoOperandsInstruction }
type DSTORE_3 struct { base.NoOperandsInstruction }

func (self *DSTORE) Execute(frame *rtda.Frame)  {
	_dstore(frame, self.Index)
}

func (self *DSTORE_0) Execute(frame *rtda.Frame)  {
	_dstore(frame, 0)
}

func (self *DSTORE_1) Execute(frame *rtda.Frame)  {
	_dstore(frame, 1)
}

func (self *DSTORE_2) Execute(frame *rtda.Frame)  {
	_dstore(frame, 2)
}

func (self *DSTORE_3) Execute(frame *rtda.Frame)  {
	_dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}