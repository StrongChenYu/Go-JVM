package heap

//方法描述符
type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

//其实就是不以2^n的倍数扩充数组
//而是以4*n的倍数扩充数组
func (self *MethodDescriptor) addParameterType(param string) {
	length := len(self.parameterTypes)
	if length == cap(self.parameterTypes) {
		//扩容
		temp := make([]string, length, length+4)
		copy(temp, self.parameterTypes)
		self.parameterTypes = temp
	}

	self.parameterTypes = append(self.parameterTypes, param)
}

func (self *MethodDescriptor) ParameterTypes() []string {
	return self.parameterTypes
}
