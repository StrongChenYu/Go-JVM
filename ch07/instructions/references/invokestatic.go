package references

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
	"go-jvm/ch07/rtda/heap"
)

type INVOKE_STATIC struct { base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame)  {
	curCp := frame.Method().Class().ConstantPool()
	methodRef := curCp.GetConstant(self.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()

	class := method.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, method)
}