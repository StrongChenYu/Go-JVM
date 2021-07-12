package classfile

type ConstantNameAndTypeInfo struct {
	//名称
	nameIndex uint16
	//描述
	typeIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUnit16()
	c.typeIndex = reader.readUnit16()
}
