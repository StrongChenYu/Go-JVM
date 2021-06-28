package heap

import "strings"

type MethodDescriptorParser struct {
	raw				string
	offset 			int
	methodDescriptor	 	*MethodDescriptor
}


func ParseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{
		raw:              descriptor,
		offset:           0,
		methodDescriptor: &MethodDescriptor{
			parameterTypes: make([]string, 0),
			returnType:     "",
		},
	}
	return parser.parse()
}


func (self *MethodDescriptorParser) parse() *MethodDescriptor {
	self.startParam()
	self.parseParamTypes()
	self.endParam()
	self.parseReturnTypes()
	self.finish()
	return self.methodDescriptor
}

func (self *MethodDescriptorParser) startParam() {
	if self.readUint8() != '(' {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) parseParamTypes() {
	for {
		t := self.parseFieldType()
		if t != "" {
			self.methodDescriptor.addParameterType(t)
		} else {
			break
		}
	}
}


func (self *MethodDescriptorParser) parseFieldType() string {
	descriptor := self.readUint8()
	switch descriptor {
	case 'B':
		//byte
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		//boolean
		return "Z"
	case 'L':
		//引用
		return self.parseObjectType()
	case '[':
		//数组
		return self.parseArrayType()
	default:
		self.unreadUint8()
		return ""
	}
}

//解析对象
func (self *MethodDescriptorParser) parseObjectType() string {
	unreadPart := self.raw[self.offset:]
	idx := strings.IndexRune(unreadPart, ';')

	if idx == -1 {
		self.causePanic()
		return ""
	} else {
		//Ljava/lang/String;
		//idx = 17
		//idx + 1 = len
		//start + len = end

		start := self.offset - 1
		end := self.offset + idx + 1

		descriptor := self.raw[start : end]
		self.offset = end

		return descriptor
	}
}


//解析数组类型
func (self *MethodDescriptorParser) parseArrayType() string {
	start := self.offset - 1
	self.parseFieldType()
	end := self.offset
	return self.raw[start : end]
}


func (self *MethodDescriptorParser) endParam() {
	if self.readUint8() != ')' {
		self.causePanic()
	}
}


func (self *MethodDescriptorParser) parseReturnTypes() {
	if self.readUint8() == 'V' {
		self.methodDescriptor.returnType = "V"
		return
	}

	self.unreadUint8()
	descriptor := self.parseFieldType()

	if descriptor != "" {
		self.methodDescriptor.returnType = descriptor
		return
	}

	self.causePanic()
}

func (self *MethodDescriptorParser) finish() {
	if self.offset != len(self.raw) {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + self.raw)
}

func (self *MethodDescriptorParser) readUint8() uint8 {
	b := self.raw[self.offset]
	self.offset++
	return b
}

func (self *MethodDescriptorParser) unreadUint8() {
	self.offset--
}





