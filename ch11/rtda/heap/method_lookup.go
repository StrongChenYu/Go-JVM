package heap

func LookUpMethodInClass(c *Class, name string, descriptor string) *Method {
	for class := c; class != nil; class = class.superClass {
		methods := class.Methods()
		for _, method := range methods {
			if method.Descriptor() == descriptor && method.Name() == name {
				return method
			}
		}
	}
	return nil
}

func LookUpMethodInInterface(inters []*Class, name string, descriptor string) *Method {
	for _, inter := range inters {
		//到接口之中去找
		for _, method := range inter.methods {
			if method.Descriptor() == descriptor && method.Name() == name {
				return method
			}
		}
		//java中支持多继承，所以是去interfaces数组中去寻找
		//在去继承的接口中寻找
		method := LookUpMethodInInterface(inters, name, descriptor)

		if method != nil {
			return nil
		}
	}

	return nil
}
