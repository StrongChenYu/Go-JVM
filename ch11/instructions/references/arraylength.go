package references

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrayRef := stack.PopRef()

	if arrayRef == nil {
		panic("java.lang.NullPointException")
	}

	arrLen := arrayRef.ArrayLength()
	stack.PushInt(arrLen)
}
