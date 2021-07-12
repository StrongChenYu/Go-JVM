package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

func init() {
	//public native String intern();
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
