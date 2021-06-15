package stores

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

//将操作数栈顶Long类型数据放到位置为i的LocalVars
type ASTORE struct { base.Index8Instruction }
type ASTORE_0 struct { base.NoOperandsInstruction }
type ASTORE_1 struct { base.NoOperandsInstruction }
type ASTORE_2 struct { base.NoOperandsInstruction }
type ASTORE_3 struct { base.NoOperandsInstruction }

func (self *ASTORE) Execute(frame *rtda.Frame)  {
	_astore(frame, self.Index)
}

func (self *ASTORE_0) Execute(frame *rtda.Frame)  {
	_astore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *rtda.Frame)  {
	_astore(frame, 1)
}

func (self *ASTORE_2) Execute(frame *rtda.Frame)  {
	_astore(frame, 2)
}

func (self *ASTORE_3) Execute(frame *rtda.Frame)  {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}