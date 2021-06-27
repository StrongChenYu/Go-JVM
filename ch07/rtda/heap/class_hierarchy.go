package heap

//check operation follow
//self a = instance (other type)
func (self *Class) IsAssignableFrom(other *Class) bool {

	//判断是否相等
	if other == self {
		//如果为true，则判断成功
		return true
	}

	if self.IsInterface() {
		return other.IsImplements(self)
	} else {
		return other.IsSubClassOf(self)
	}

}


//check operation follow
//other a = instance (self type)
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}


//check operation follow
//inter a = instance (self type)
func (self *Class) IsImplements (cfInter *Class) bool  {
	for c := self; c != nil; c = c.superClass {
		for _, inter  := range c.interfaces {
			if inter == cfInter || inter.isSubInterfaceOf(cfInter) {
				return true
			}
		}
	}
	return false;
}

//check operation follow
//inter a = instance (self type)
func (self *Class) isSubInterfaceOf(inter *Class) bool {
	for _, i := range self.interfaces {
		if i == inter || i.isSubInterfaceOf(inter) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(sonClass *Class) bool {
	return sonClass.IsSubClassOf(self)
}