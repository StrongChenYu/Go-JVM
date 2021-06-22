package control

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

//goto是无条件跳转
type GOTO struct { base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.Offset)
}