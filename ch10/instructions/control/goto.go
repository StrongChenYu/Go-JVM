package control

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
)

//goto是无条件跳转
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
