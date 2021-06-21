package heap

import "math"

type Slot struct {
	num			int32
	ref			*Object
}


//局部变量表
type Slots []Slot

//new一个field槽
func newSlots(fieldCount uint) Slots {
	if fieldCount > 0 {
		return make([]Slot, fieldCount)
	}
	return nil
}

//设置index位置为int值
func (self Slots) SetInt(index uint, val int32)  {
	self[index].num = val
}

//获取index处的int值
func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

//设置float值
func (self Slots) SetFloat(index uint, val float32)  {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

//返回float值
func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

//long类型数据
func (self Slots) SetLong(index uint, val int64)  {
	//低32位
	self[index].num = int32(val)
	//高32位
	self[index + 1].num = int32(val >> 32)
}

//获取long类型数据
func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index + 1].num)
	//这里不用uint64来计算
	return int64(high) << 32 | int64(low)
}

//获取double数据
func (self Slots) SetDouble(index uint, val float64)  {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

//设置double数据
func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}


//设置引用值
func (self Slots) SetRef(index uint, ref *Object)  {
	self[index].ref = ref
}

//获取引用值
func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}
