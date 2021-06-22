package comparisons

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

//浮点数比较
//两个指令的区别在于，对于NAN的结果定义不同
//Otherwise, at least one of value1' or value2' is NaN.
//The fcmpg instruction pushes the int value 1 onto the operand stack and the fcmpl instruction pushes the int value -1 onto the operand stack.
type FCMPG struct { base.NoOperandsInstruction }
type FCMPL struct { base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame)  {
	_fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rtda.Frame)  {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	op2 := stack.PopFloat()
	op1 := stack.PopFloat()

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