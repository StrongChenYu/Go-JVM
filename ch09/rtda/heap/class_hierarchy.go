package heap

//check operation follow
//self a = instance (other type)
func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.IsJlObject()
			} else {
				// t is interface
				return t.IsSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.IsJlObject()
			} else {
				// t is interface
				return t.IsJlCloneable() || t.IsJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
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
func (self *Class) IsImplements(cfInter *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, inter := range c.interfaces {
			if inter == cfInter || inter.isSubInterfaceOf(cfInter) {
				return true
			}
		}
	}
	return false
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

func (self *Class) IsSuperInterfaceOf(s *Class) bool {
	return s.isSubInterfaceOf(self)
}
