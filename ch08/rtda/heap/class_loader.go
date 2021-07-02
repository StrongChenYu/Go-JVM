package heap

import (
	"fmt"
	"go-jvm/ch08/classfile"
	"go-jvm/ch08/classpath"
)

//一个简单的类加载器
type ClassLoader struct {
	cp          *classpath.Classpath
	classMap    map[string]*Class
	verboseFlag bool
}

//构造函数
func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}

	if name[0] == '[' {
		return self.loadArrayClass(name)
	}

	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //todo
		name:        name,
		loader:      self,
		superClass:  self.LoadClass("java/lang/Object"),
		initStarted: true,
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

//加载类
//loadClass -> loadNonArrayClass -> readClass(read byte data)
//						|
//						----------> defineClass(byte data ---> Class)
//						|
//						----------> link------> verify (do nothing)
//										|
//										------> prepare (cal field index in slots (include static, and init static field))
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)

	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func link(class *Class) {
	verify(class)
	prepare(class)
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
	//!!!!!!!!!!!!!!!!!!!!!!!
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

func verify(class *Class) {
	//DO NOTHING
	//see java specification 4.10
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotFields(class)
	allocAndInitStaticVars(class)
}

//type Object struct {
//	//用于存放引用类型
//	class 		*Class
//	fields 		Slots
//}
//这部分计算的是，field中的值，在object中的fields中的index
//Java中的静态变量都存储在类中，非静态变量都存储在实例化的Object中
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		//考虑来自父类
		slotId = class.superClass.instanceSlotCount
	}

	//找到类中的每一个非静态
	//然后根据这个非静态属性
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

//这部分计算的是java类中的静态变量索引
//因为静态变量是属于类的，所以即存储在class.staticVars中
func calcStaticFieldSlotFields(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//然后给静态变量初始化
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
