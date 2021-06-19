package heap

import "go-jvm/ch06/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method 		*Method
}

func newInterfaceMethodRef(cp *ConstantPool, info *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
