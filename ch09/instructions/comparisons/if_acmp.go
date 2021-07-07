package comparisons

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
)

//根据引用相同是否跳转
type IF_ACMPEQ struct{ base.BranchInstruction }
type IF_ACMPNE struct{ base.BranchInstruction }

//如果相同就跳转
func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopRef()
	op1 := stack.PopRef()

	if op1 == op2 {
		base.Branch(frame, self.Offset)
	}
}

//如果不同就跳转
func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopRef()
	op1 := stack.PopRef()

	if op1 != op2 {
		base.Branch(frame, self.Offset)
	}
}
