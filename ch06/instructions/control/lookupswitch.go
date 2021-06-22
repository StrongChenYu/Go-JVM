package control

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset 		int32
	npairs 				int32
	matchOffsets 		[]int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs * 2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i + 1]
			base.Branch(frame, int(offset))
			return
		}
	}

	//默认跳转
	base.Branch(frame, int(self.defaultOffset))
}



