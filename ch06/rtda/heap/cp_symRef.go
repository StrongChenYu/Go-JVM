package heap

//注意：ClassRef并没有给class指针赋值，解析的目的正在于此
//解析的目的是解析class
//目的是cp.class -> class
type SymRef struct {
	cp 			*ConstantPool
	className 	string
	class 		*Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef()  {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}