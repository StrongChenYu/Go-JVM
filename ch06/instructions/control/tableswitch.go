package control

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)
//??????? go语言里面int到底是几位？？？

type TABLE_SWITCH struct {
	defaultOffset 		int32
	low 				int32
	high				int32
	jumpOffsets 		[]int32
}

//jumpOffsets中存储着跳转的地址数组
func (self *TABLE_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	//内存对齐
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	self.jumpOffsets = reader.ReadInt32s(self.high - self.low + 1)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	//默认跳转指令
	base.Branch(frame, offset)
}



