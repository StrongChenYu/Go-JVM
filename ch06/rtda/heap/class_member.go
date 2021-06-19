package heap

import "go-jvm/ch06/classfile"

//field和method的父类，抽象出来
type ClassMember struct {
	accessFlags			uint16
	name 				string
	descriptor 			string
	class 				*Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo)  {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
