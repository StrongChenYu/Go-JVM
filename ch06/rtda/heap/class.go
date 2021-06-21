package heap

import (
	"go-jvm/ch06/classfile"
	"strings"
)

type Class struct {
	accessFlags 			uint16
	name 					string
	superClassName 			string
	interfaceNames			[]string
	constantPool 			*ConstantPool
	fields 					[]*Field
	methods 				[]*Method
	loader 					*ClassLoader
	superClass				*Class
	interfaces 				[]*Class
	instanceSlotCount       uint
	staticSlotCount			uint
	staticVars 				Slots
}


func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{
		accessFlags:    cf.AccessFlags(),
		name:           cf.ClassName(),
		superClassName: cf.SuperClassName(),
		interfaceNames: cf.InterfaceNames(),
	}

	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}


func (self *Class) IsAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) IsPublic() bool {
	return self.accessFlags & ACC_PUBLIC != 0
}

func (self *Class) IsInterface() bool {
	return self.accessFlags & ACC_INTERFACE != 0
}

func (self *Class) IsAbstract() bool {
	return self.accessFlags & ACC_ABSTRACT != 0
}

//input : java/lang/Object
//output: java/lang
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}


func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}