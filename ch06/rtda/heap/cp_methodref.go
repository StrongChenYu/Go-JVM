package heap

import "go-jvm/ch06/classfile"

type MethodRef struct {
	MemberRef
	method 		*Method
}

func newMethodRef(cp *ConstantPool, info *classfile.ConstantMethodRefInfo) *MemberRef {
	ref := &MemberRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
