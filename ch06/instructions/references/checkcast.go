package references

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
	"go-jvm/ch06/rtda/heap"
)

//判断operandStack中弹出的ref，是不是base.index16Instruction指向的常量池classref
//与instanceof不同的是，会直接抛异常
type CHECK_CAST struct { base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()

	ref := stack.PopRef()

	//null can be casted to any type
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if !ref.IsInstanceOf(class) {
		//throw exception
		panic("java.lang.ClassCastException")
	}
	stack.PushRef(ref)
}