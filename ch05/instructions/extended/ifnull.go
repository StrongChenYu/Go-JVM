package extended

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type IFNULL struct { base.BranchInstruction }
type IFNONULL struct { base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONULL) Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}