package rtda

import "go-jvm/ch08/rtda/heap"

type Thread struct {
	//Thread的PC指针
	pc int
	//线程栈
	stack *Stack
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

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

//thread.newFrame 将会产生一个新的栈帧，这个栈帧可以直接由方法产生
//然后把这个栈帧压入栈顶
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.IsEmpty()
}
