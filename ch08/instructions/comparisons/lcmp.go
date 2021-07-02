package comparisons

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

//长整型数字
type LCMP struct{ base.NoOperandsInstruction }

//[a][b][c]->
//if b > c push 1
//if b == c push 0
//if b < c push -1
func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopLong()
	op1 := stack.PopLong()

	if op1 > op2 {
		stack.PushInt(1)
	} else if op1 < op2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
