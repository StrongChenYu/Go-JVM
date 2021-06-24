package heap

import "go-jvm/ch07/classfile"

type MethodRef struct {
	MemberRef
	method 		*Method
}

func newMethodRef(cp *ConstantPool, info *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {
	curClass := self.cp.class
	c := self.ResolvedClass()

	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(curClass) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupMethod(c *Class, name string, descriptor string) *Method {
	method := LookUpMethodInClass(c, name, descriptor)
	if method == nil {
		method = LookUpMethodInInterface(c.interfaces, name, descriptor)
	}
	return method
}