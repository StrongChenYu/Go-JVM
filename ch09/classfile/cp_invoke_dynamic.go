package classfile

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (c *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	c.descriptorIndex = reader.readUnit16()
}

type ConstantInvokeDynamicInfo struct {
	bootStrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (c *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	c.bootStrapMethodAttrIndex = reader.readUnit16()
	c.nameAndTypeIndex = reader.readUnit16()
}

type ConstantMethodHandlerInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (c *ConstantMethodHandlerInfo) readInfo(reader *ClassReader) {
	c.referenceKind = reader.readUnit8()
	c.referenceIndex = reader.readUnit16()
}
