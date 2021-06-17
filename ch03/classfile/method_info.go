package classfile

type MemberInfo struct {
	cp					ConstantPool
	accessFlags			uint16
	nameIndex			uint16
	descriptorIndex		uint16
	attributes 			[]AttributeInfo
}

func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readerMember(reader, pool)
	}
	return members
}

func readerMember(reader *ClassReader, pool ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              pool,
		accessFlags:     reader.readUnit16(),
		nameIndex:       reader.readUnit16(),
		descriptorIndex: reader.readUnit16(),
		attributes:      readAttributes(reader, pool),
	}
}

func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}