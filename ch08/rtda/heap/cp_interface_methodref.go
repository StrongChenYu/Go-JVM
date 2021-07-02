package heap

import "go-jvm/ch08/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, info *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}

func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	curClass := self.cp.class
	c := self.ResolvedClass()

	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(c, self.name, self.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(curClass) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupInterfaceMethod(c *Class, name string, descriptor string) *Method {
	for _, method := range c.methods {
		if method.descriptor == descriptor && method.name == name {
			return method
		}
	}
	return LookUpMethodInInterface(c.interfaces, name, descriptor)
}
