package heap

import "go-jvm/ch07/classfile"

type MethodRef struct {
	MemberRef
	method 		*Method
}

func newMethodRef(cp *ConstantPool, info *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
