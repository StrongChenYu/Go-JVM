package references

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
	"go-jvm/ch07/rtda/heap"
)

//判断operandStack中弹出的ref，是不是base.index16Instruction指向的常量池classref
type INSTANCE_OF struct { base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()

	ref := stack.PopRef()

	//null instanceof class always null
	if ref == nil {
		stack.PushInt(0)
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	//ref instanceof class
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

