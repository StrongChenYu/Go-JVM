package loads

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
	"go-jvm/ch10/rtda/heap"
)

type AALOAD struct{ base.NoOperandsInstruction }
type BALOAD struct{ base.NoOperandsInstruction }
type CALOAD struct{ base.NoOperandsInstruction }
type DALOAD struct{ base.NoOperandsInstruction }
type FALOAD struct{ base.NoOperandsInstruction }
type IALOAD struct{ base.NoOperandsInstruction }
type LALOAD struct{ base.NoOperandsInstruction }
type SALOAD struct{ base.NoOperandsInstruction }

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Refs()
	checkIndex(len(arrayData), index)
	stack.PushRef(arrayData[index])
}

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Bytes()
	checkIndex(len(arrayData), index)
	stack.PushInt(int32(arrayData[index]))
}

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Chars()
	checkIndex(len(arrayData), index)
	stack.PushInt(int32(arrayData[index]))
}

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Doubles()
	checkIndex(len(arrayData), index)
	stack.PushDouble(arrayData[index])
}

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Floats()
	checkIndex(len(arrayData), index)
	stack.PushFloat(arrayData[index])
}

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Ints()
	checkIndex(len(arrayData), index)
	stack.PushInt(arrayData[index])
}

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Ints()
	checkIndex(len(arrayData), index)
	stack.PushInt(arrayData[index])
}

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	arrayData := arrRef.Longs()
	checkIndex(len(arrayData), index)
	stack.PushLong(arrayData[index])
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
