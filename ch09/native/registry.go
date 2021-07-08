package native

import "go-jvm/ch09/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, descriptor string, nativeMethod NativeMethod) {
	key := className + "~" + methodName + "~" + descriptor
	registry[key] = nativeMethod
}

func FindNativeMethod(className, methodName, descriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + descriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if descriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	//doNothing
}
