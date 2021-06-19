package heap

import "go-jvm/ch06/classfile"

//字段的符号引用
type FieldRef struct {
	MemberRef
	field  *Field
}

func newFieldRef(cp *ConstantPool, info *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
