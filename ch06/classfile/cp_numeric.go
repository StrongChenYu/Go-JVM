package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) Val() int32 {
	return c.val
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	c.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) Val() float32 {
	return c.val
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	c.val = math.Float32frombits(bytes)
}

//Long类型
type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) Val() int64 {
	return c.val
}

func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit64()
	c.val = int64(bytes)
}

//Double类型
type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) Val() float64 {
	return c.val
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit64()
	c.val = math.Float64frombits(bytes)
}





