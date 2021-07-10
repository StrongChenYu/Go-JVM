package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
	"unsafe"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	//public native int hashCode();
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	//protected native Object clone() throws CloneNotSupportedException;
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")

	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	ref := this.Class().JClass()
	frame.OperandStack().PushRef(ref)
}
