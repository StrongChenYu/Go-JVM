package rtda

import (
	"go-jvm/ch09/rtda/heap"
	"math"
)

//局部变量表
type LocalVars []Slot

//new一个局部变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//设置index位置为int值
func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}

//获取index处的int值
func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}

//设置float值
func (lv LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	lv[index].num = int32(bits)
}

//返回float值
func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].num)
	return math.Float32frombits(bits)
}

//long类型数据
func (lv LocalVars) SetLong(index uint, val int64) {
	//低32位
	lv[index].num = int32(val)
	//高32位
	lv[index+1].num = int32(val >> 32)
}

//获取long类型数据
func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	//这里不用uint64来计算
	return int64(high)<<32 | int64(low)
}

//获取double数据
func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

//设置double数据
func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

//设置引用值
func (lv LocalVars) SetRef(index uint, ref *heap.Object) {
	lv[index].ref = ref
}

//获取引用值
func (lv LocalVars) GetRef(index uint) *heap.Object {
	return lv[index].ref
}

func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}

func (self LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}
