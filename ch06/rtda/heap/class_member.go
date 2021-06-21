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

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) IsPublic() bool {
	return self.accessFlags & ACC_PUBLIC != 0
}

func (self *ClassMember) IsPrivate() bool {
	return self.accessFlags & ACC_PRIVATE != 0
}

func (self *ClassMember) IsProtected() bool {
	return self.accessFlags & ACC_PROTECTED != 0
}

//判断d是否可以访问self
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}

	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}

	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}

	return d == c
}

func (self *ClassMember) Class() *Class {
	return self.class
}


func (self *ClassMember) Name() string {
	return self.name
}