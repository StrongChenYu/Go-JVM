package rtda

//栈帧
//可以理解为一个函数的调用栈帧
type Frame struct {
	lower			*Frame
	localVars 		LocalVars
	operandStack    *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

//get方法
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

//get方法
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}