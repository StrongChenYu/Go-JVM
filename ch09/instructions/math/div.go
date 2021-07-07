package math

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"math"
)

type IDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }
type DDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopInt()
	op1 := stack.PopInt()

	if op2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := op1 / op2
	stack.PushInt(result)
}

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopLong()
	op1 := stack.PopLong()

	if op2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := op1 / op2
	stack.PushLong(result)
}

//取余运算可以为0，但是/运算符绝对不能为0
func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopDouble()
	op1 := stack.PopDouble()

	//双精度浮点数判断是否为0
	if math.Abs(op2) <= 1e-6 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := op1 / op2
	stack.PushDouble(result)
}

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	op2 := stack.PopFloat()
	op1 := stack.PopFloat()

	//单精度浮点数判断是否为0
	if math.Abs(float64(op2)) <= 1e-15 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := op1 / op2
	stack.PushFloat(result)
}
