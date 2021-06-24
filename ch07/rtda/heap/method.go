package heap

import "go-jvm/ch07/classfile"

//描述方法的
type Method struct {
	ClassMember
	maxStack 		uint
	maxLocal 		uint
	code 			[]byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))

	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].copyCodeAttributes(cfMethod)
		methods[i].copyMemberInfo(cfMethod)
		methods[i].class = class
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

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) MaxLocal() uint {
	return self.maxLocal
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}