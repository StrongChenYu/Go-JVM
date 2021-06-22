package base

import "go-jvm/ch06/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {

}

func (n *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
}


func (n *NoOperandsInstruction) Execute(frame *rtda.Frame) {

}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *BranchInstruction) Execute(frame *rtda.Frame) {

}

type Index8Instruction struct {
	Index 		uint
}

func (self *Index8Instruction) Execute(frame *rtda.Frame) {
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader)  {
	self.Index = uint(reader.ReadUint8())
}


type Index16Instruction struct {
	Index 		uint
}

func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}

func (self Index16Instruction) Execute(frame *rtda.Frame) {
}
