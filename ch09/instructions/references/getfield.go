package references

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

type GET_FIELD struct{ base.Index16Instruction }

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()

	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := stack.PopRef()

	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlodId()
	slots := ref.Fields()

	if slots == nil {
		panic("java.lang.NullPointerException")
	}

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
