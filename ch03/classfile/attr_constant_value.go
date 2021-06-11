package classfile

/**
	常量值的索引位置，
	这个索引可以指向Constant_Integer,Constant_Long等一系列
 */
type ConstantValueAttribute struct {
	constantValueIdx 		uint16
}

func (c *ConstantValueAttribute) readInfo(reader *ClassReader) {
	c.constantValueIdx = reader.readUnit16()
}

func (c *ConstantValueAttribute) ConstantValueIdx() uint16 {
	return c.constantValueIdx
}

