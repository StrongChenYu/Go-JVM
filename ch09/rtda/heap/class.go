package heap

import (
	"go-jvm/ch09/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots

	//判断有没有执行cint方法
	initStarted bool

	//指向类对象
	jClass *Object
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
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}

func (self *Class) IsInterface() bool {
	return self.accessFlags&ACC_INTERFACE != 0
}

func (self *Class) IsAbstract() bool {
	return self.accessFlags&ACC_ABSTRACT != 0
}

func (self *Class) IsSuper() bool {
	return self.accessFlags&ACC_SUPER != 0
}

//input : java/lang/Object
//output: java/lang
func (self *Class) GetPackageName() string {
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
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.Descriptor() == descriptor && method.Name() == name {
			return method
		}
	}
	return nil
}

func (self *Class) GetclinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

//arrayClass
//java/lang/Object => [java/lang/Object
//int => [i
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) IsJlObject() bool {
	return self.name == "java/lang/Obejct"
}

func (self *Class) IsJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

func (self *Class) IsJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) getField(name string, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
