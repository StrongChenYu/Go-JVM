package heap

type Object struct {
	//用于存放引用类型
	class *Class
	data  interface{}
	extra interface{}
}

func (o *Object) Extra() interface{} {
	return o.extra
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(o.class)
}

func (self *Object) SetRefVar(name, descriptor string, chars *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.SlodId(), chars)
}

func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (self *Object) Clone() *Object {
	return &Object{
		class: self.class,
		data:  self.cloneDate(),
		extra: self.extra,
	}
}

func (self *Object) cloneDate() interface{} {
	switch self.data.(type) {
	case []int8:
		oldData := self.data.([]int8)
		newData := make([]int8, len(oldData))
		copy(newData, oldData)
		return newData
	case []int16:
		oldData := self.data.([]int16)
		newData := make([]int16, len(oldData))
		copy(newData, oldData)
		return newData
	case []int32:
		oldData := self.data.([]int32)
		newData := make([]int32, len(oldData))
		copy(newData, oldData)
		return newData
	case []int64:
		oldData := self.data.([]int64)
		newData := make([]int64, len(oldData))
		copy(newData, oldData)
		return newData
	case []uint16:
		oldData := self.data.([]uint16)
		newData := make([]uint16, len(oldData))
		copy(newData, oldData)
		return newData
	case []float32:
		oldData := self.data.([]float32)
		newData := make([]float32, len(oldData))
		copy(newData, oldData)
		return newData
	case []float64:
		oldData := self.data.([]float64)
		newData := make([]float64, len(oldData))
		copy(newData, oldData)
		return newData
	case []*Object:
		oldData := self.data.([]*Object)
		newData := make([]*Object, len(oldData))
		copy(newData, oldData)
		return newData
	default:
		oldSlots := self.data.(Slots)
		_newSlots := newSlots(uint(len(oldSlots)))
		copy(_newSlots, oldSlots)
		return _newSlots
	}
}
