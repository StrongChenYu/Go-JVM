package references

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	classRef := frame.Method().Class().ConstantPool().GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	count := stack.PopInt()

	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := class.ArrayClass()
	ref := arrClass.NewArray(uint(count))
	stack.PushRef(ref)
}
