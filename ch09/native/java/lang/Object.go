package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Class", "getPrimitiveClass", "[Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushBoolean(false)
}

func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.JStrToGoStr(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	ref := this.Class().JClass()
	frame.OperandStack().PushRef(ref)
}
