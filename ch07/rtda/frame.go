package rtda

import "go-jvm/ch07/rtda/heap"

//栈帧
//可以理解为一个函数的调用栈帧
//每一个函数都有一个本地变量表和操作数栈
type Frame struct {
	lower			*Frame
	localVars 		LocalVars
	operandStack    *OperandStack
	thread 			*Thread
	method 			*heap.Method
	nextPC 			int
}



func (self *Frame) NextPC() int {
	return self.nextPC
}

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread: thread,
		method: method,
		localVars:    newLocalVars(method.MaxLocal()),
		operandStack: newOperandStack(method.MaxStack()),
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

//get方法
func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func (self *Frame) Method() *heap.Method {
	return self.method
}