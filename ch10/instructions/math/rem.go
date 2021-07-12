package math

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
	"math"
)

//取余运算
//...[a][b][c]->
// result = b % c
type DREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	//op2代表计算中顺序为2
	op2 := stack.PopInt()
	//op1代表计算中顺序为1
	op1 := stack.PopInt()

	if op2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := op1 % op2
	stack.PushInt(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	op2 := stack.PopLong()
	//op1代表计算中顺序为1
	op1 := stack.PopLong()

	if op2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := op1 % op2
	stack.PushLong(result)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	op2 := stack.PopDouble()
	op1 := stack.PopDouble()

	result := math.Mod(op1, op2)
	stack.PushDouble(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	//转化成64位
	op2 := float64(stack.PopFloat())
	op1 := float64(stack.PopFloat())

	//???????这样会不会丢失精度
	result := float32(math.Mod(op1, op2))
	stack.PushFloat(result)
}
