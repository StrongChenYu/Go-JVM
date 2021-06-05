package classfile

type MemberInfo struct {
	cp					ConstantPool
	accessFlags			uint16
	nameIndex			uint16
	descriptorIndex		uint16
	attributes 			[]AttributeInfo
}

func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {
	return nil
}