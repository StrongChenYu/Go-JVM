package heap

import "go-jvm/ch07/classfile"

//描述方法的
type Method struct {
	ClassMember
	maxStack 		uint
	maxLocal 		uint
	code 			[]byte
	argSlotCount	int
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))

	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyCodeAttributes(cfMethod)
		methods[i].copyMemberInfo(cfMethod)
		methods[i].calcArgSlotCount()
	}

	return methods
}

func (self *Method) copyCodeAttributes(cfMethod *classfile.MemberInfo)  {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = uint(codeAttr.MaxStack())
		self.maxLocal = uint(codeAttr.MaxLocals())
		self.code = codeAttr.Code()
	}
}

func (self *Method) calcArgSlotCount() {
	methodDescriptor := ParseMethodDescriptor(self.descriptor)
	for _, paramType := range methodDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "L" || paramType == "D" {
			self.argSlotCount++
		}
	}

	//如果方法不是静态的，函数还有一个变量（个人觉得应该是保存自引用）
	if !self.IsStatic() {
		self.argSlotCount++
	}
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
	return self.accessFlags & ACC_STATIC != 0
}

func (self *Method) IsAbstract() bool {
	return self.accessFlags & ACC_ABSTRACT != 0
}

func (self *Method) IsNative() bool {
	return self.accessFlags & ACC_NATIVE != 0
}