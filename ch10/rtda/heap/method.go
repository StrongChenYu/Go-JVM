package heap

import "go-jvm/ch10/classfile"

//描述方法的
type Method struct {
	ClassMember
	maxStack        uint
	maxLocal        uint
	code            []byte
	argSlotCount    int
	exceptionTable  ExceptionTable
	lineNumberTable *classfile.LineNumberTableAttribute
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))

	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}

	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}

	method = &Method{}
	method.class = class
	method.copyCodeAttributes(cfMethod)
	method.copyMemberInfo(cfMethod)
	md := ParseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)

	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}

	return method
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocal = uint(self.argSlotCount)

	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1}
	case 'D':
		self.code = []byte{0xfe, 0xaf}
	case 'F':
		self.code = []byte{0xfe, 0xae}
	case 'J':
		self.code = []byte{0xfe, 0xad}
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0}
	default:
		self.code = []byte{0xfe, 0xac}
	}
}

func (self *Method) copyCodeAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = uint(codeAttr.MaxStack())
		self.maxLocal = uint(codeAttr.MaxLocals())
		self.code = codeAttr.Code()
		self.lineNumberTable = codeAttr.LineNumberAttribute()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), self.class.constantPool)
	}
}

func (self *Method) calcArgSlotCount(parameterTypes []string) {
	for _, paramType := range parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}

	//如果方法不是静态的，函数还有一个变量（个人觉得应该是保存自引用）
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) MaxLocal() uint {
	return self.maxLocal
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) ArgSlotCount() int {
	return self.argSlotCount
}

func (self *Method) IsStatic() bool {
	return self.accessFlags&ACC_STATIC != 0
}

func (self *Method) IsAbstract() bool {
	return self.accessFlags&ACC_ABSTRACT != 0
}

func (self *Method) IsNative() bool {
	return self.accessFlags&ACC_NATIVE != 0
}

func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}

	if self.lineNumberTable == nil {
		return -1
	}

	return self.lineNumberTable.GetLineNumber(pc)
}
