package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	ref := this.Class().JClass()
	frame.OperandStack().PushRef(ref)
}
