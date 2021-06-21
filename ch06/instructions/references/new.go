package references

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
	"go-jvm/ch06/rtda/heap"
)

//new Object()这个Object就是一个符号引用
type NEW struct { base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame)  {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
