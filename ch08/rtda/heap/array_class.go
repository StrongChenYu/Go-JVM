package heap

func (self *Class) NewArray(u uint) *Object {
	if !self.IsArray() {
		panic("Not array class: " + self.name)
	}
}
