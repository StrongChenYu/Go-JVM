package references

import (
	"fmt"
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()

	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//方法第一个参数就是指向实例的对象
	ref := frame.OperandStack().GetRefFromTop(uint(method.ArgSlotCount() - 1))

	if ref == nil {
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointException")
	}

	//see invokespecial.go find why
	if method.IsProtected() &&
		method.Class().IsSuperClassOf(currentClass) &&
		method.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	//在ref中指向的实例代表的类的方法
	methodToBeInvoked := heap.LookUpMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.JStrToGoStr(jStr)
		fmt.Println(goStr)
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()

}
