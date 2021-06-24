package references

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
	"go-jvm/ch07/rtda/heap"
)


//获取某个变量的静态值，然后推入到操作数栈中
type GET_STATIC struct { base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()

	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	staticVars := field.Class().StaticVars()
	descriptor := field.Descriptor()
	slotId := field.SlodId()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := staticVars.GetInt(slotId)
		stack.PushInt(val)
	case 'J':
		val := staticVars.GetLong(slotId)
		stack.PushLong(val)
	case 'F':
		val := staticVars.GetFloat(slotId)
		stack.PushFloat(val)
	case 'D':
		val := staticVars.GetDouble(slotId)
		stack.PushDouble(val)
	case 'L', '[':
		val := staticVars.GetRef(slotId)
		stack.PushRef(val)
	}
}
