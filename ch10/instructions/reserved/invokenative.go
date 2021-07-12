package reserved

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/native"
	"go-jvm/ch10/rtda"
)

import _ "go-jvm/ch10/native/java/lang"
import _ "go-jvm/ch10/native/sun/misc"

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
