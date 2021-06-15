package rtda

import "math"

//操作栈
type OperandStack struct {
	size 		uint
	slots 		[]Slot
}

//操作数栈
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

//size指向栈顶
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32)  {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

//slot[low][high]
func (self *OperandStack) PushLong(val int64) {
	//低32位
	self.slots[self.size].num = int32(val)
	//高32位
	self.slots[self.size + 1].num = int32(val >> 32)
	self.size += 2
}

//slot[low][high]
func (self *OperandStack) PopLong() int64 {
	self.size--
	highBits := uint32(self.slots[self.size].num)
	lowBits := uint32(self.slots[self.size - 1].num)
	return int64(highBits) << 32 | int64(lowBits)
}

//slot[low][high]
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

//pop出double数值
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *Object)  {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	//这里要置nil，不然会妨碍后面操作
	self.slots[self.size].ref = nil
	return ref
}

//为了实现栈指令
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

//获取操作数栈的大小
func (self *OperandStack) Size() uint {
	return self.size
}

