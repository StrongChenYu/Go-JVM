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
