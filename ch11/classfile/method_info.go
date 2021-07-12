package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
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

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}

func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

//?????????attrInfo.(type) 这种语法到底是什么含义
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

//Descriptor
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}

//如果field实现被定义好，那么返回存储在常量池中的索引
func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
