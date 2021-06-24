package extended

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

//GOTO instruction
//self.Offset = int(reader.ReadInt16())
type GOTO_W struct { offset int }

func (self *GOTO_W) FetchOperands(reader *base.ByteCodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}

