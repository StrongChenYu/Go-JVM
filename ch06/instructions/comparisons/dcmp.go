package comparisons

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

//和Float值一样
type DCMPG struct { base.NoOperandsInstruction }
type DCMPL struct { base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame)  {
	_dcmp(frame, true)
}

func (self *DCMPL) Execute(frame *rtda.Frame)  {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	op2 := stack.PopDouble()
	op1 := stack.PopDouble()

	//NaN和任何数字的>,<,=都是false
	if op1 > op2 {
		stack.PushInt(1)
	} else if op1 < op2 {
		stack.PushInt(-1)
	} else if op1 == op2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
