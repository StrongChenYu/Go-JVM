package references

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
	"go-jvm/ch07/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index 		uint
	//count 	uint8
	//zero		uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.ByteCodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	method := methodRef.ResolvedInterfaceMethod()

	if method.IsStatic() || method.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(method.ArgSlotCount() - 1)

	if ref == nil {
		panic("java.lang.NullPointException")
	}

	methodToBeInvoked := heap.LookUpMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
