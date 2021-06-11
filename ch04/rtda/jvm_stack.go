package rtda

type Stack struct {
	maxSize 	uint
	size    	uint
	//frame链表大小
	_top		*Frame
}

//生成栈
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		//栈帧溢出
		panic("java.lang.StackOverFlowError")
	}
	//放置到栈顶
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	frame := self._top

	//这里主要是出栈
	self._top = frame.lower
	//***
	frame.lower = nil
	self.size--

	return frame
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}