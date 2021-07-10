package reserved

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
)

import _ "go-jvm/ch09/native/java/lang"
import _ "go-jvm/ch09/native/sun/misc"

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	descriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, descriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "." + descriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
