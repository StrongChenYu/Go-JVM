package rtda

import "go-jvm/ch07/rtda/heap"

type Thread struct {
	//Thread的PC指针
	pc 		int
	//线程栈
	stack   *Stack
}

func NewThread() *Thread {
	return &Thread{
		//这里默认1024个栈帧
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int)  {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame)  {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}