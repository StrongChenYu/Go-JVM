package constants

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
	"go-jvm/ch10/rtda/heap"
)

//将常量池中的数据加载到操作数栈中
type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	constant := cp.GetConstant(self.Index)

	switch constant.(type) {
	case int64:
		stack.PushLong(constant.(int64))
	case float64:
		stack.PushDouble(constant.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}

}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	constant := cp.GetConstant(index)
	class := frame.Method().Class()
	loader := class.Loader()

	switch constant.(type) {
	case int32:
		stack.PushInt(constant.(int32))
	case float32:
		stack.PushFloat(constant.(float32))
	case string:
		internedStr := heap.JString(loader, constant.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef := constant.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	default:
	}
}
