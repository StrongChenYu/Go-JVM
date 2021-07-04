package heap

type Object struct {
	//用于存放引用类型
	class *Class
	data  interface{}
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

func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.data.([]int8)))
	case []int16:
		return int32(len(self.data.([]int16)))
	case []int32:
		return int32(len(self.data.([]int32)))
	case []int64:
		return int32(len(self.data.([]int64)))
	case []uint16:
		return int32(len(self.data.([]uint16)))
	case []float32:
		return int32(len(self.data.([]float32)))
	case []float64:
		return int32(len(self.data.([]float64)))
	case []*Object:
		return int32(len(self.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
