package references

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
	"go-jvm/ch10/rtda/heap"
)

//Invoke instance method; special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	class := methodRef.ResolvedClass()
	method := methodRef.ResolvedMethod()

	//如果是构造方法，必须是方法自己的类
	if method.Name() == "<init>" && method.Class() != class {
		//this.fatherMethod();
		panic("java.lang.NoSuchMethodError")
	}

	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(uint(method.ArgSlotCount() - 1))

	if ref == nil {
		panic("java.lang.NullPointException")
	}

	//what is this mean?
	//situation One:
	//public void sonMethod() {
	//	Father father = new Father();
	//	father.fatherMethod();
	//	this.fatherMethod();
	//}
	//sonMethod中实例化父类，然后调用父类的protected方法，权限不够

	//situation two:
	//public void grandFather() {
	//	Father father = new Father();
	//	father.fatherMethod();
	//}
	//father的父类grandFather中某方法实例化father,然后调用father的方法，是可以调用的

	//这个逻辑判断的就是situation One会不会报错，前提是不在同一个包下面
	if method.IsProtected() &&
		method.Class().IsSuperClassOf(currentClass) &&
		method.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := method
	if currentClass.IsSuper() &&
		class.IsSubClassOf(currentClass) &&
		method.Name() != "<init>" {

		//去父类中查找，不在本类中查找
		methodToBeInvoked = heap.LookUpMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
