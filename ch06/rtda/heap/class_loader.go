package heap

import (
	"go-jvm/ch06/classfile"
	"go-jvm/ch06/classpath"
)

//一个简单的类加载器
type ClassLoader struct {
	cp 			*classpath.Classpath
	classMap 	map[string]*Class
}

//构造函数
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

//加载类
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	return nil
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}

	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)

	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			//??????????这里接口的类加载器，使用的和类本身一样吗？
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}


func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//DO NOTHING
	//see java specification 4.10
}

func prepare(class *Class) {
	calcInstanceFieldSlotFields(class)
	calcStaticFieldSlotFields(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotFields(class *Class)  {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slodId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotFields(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slodId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstsValueIndex()
	slotId := field.SlodId()


	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String":
			panic("todo")
		}
	}
}