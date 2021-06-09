package classfile

type ConstantValueAttribute struct {
	constantValueIdx 		uint16
}

func (c *ConstantValueAttribute) readInfo(reader *ClassReader) {
	c.constantValueIdx = reader.readUnit16()
}

func (c *ConstantValueAttribute) ConstantValueIdx() uint16 {
	return c.constantValueIdx
}

