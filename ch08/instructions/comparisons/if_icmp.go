package comparisons

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

//将栈顶的两个值弹出，然后比较，如果满足条件就跳转
type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op1 == op2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op1 != op2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op1 < op2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op1 <= op2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op1 > op2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op1 >= op2 {
		base.Branch(frame, self.Offset)
	}
}
