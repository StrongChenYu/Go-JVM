package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//读取attributeInfo
//先读取一个attribute的length
//然后返回attribute
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	length := reader.readUnit16()
	attributes := make([]AttributeInfo, length)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}

/**
attribute_info {
	1. name_idx //根据这个idx判断attribute的类型
	2. length // attribute中的内容长度
	3. u1 长度内容
}
*/
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	//读取attribute的类型
	attrNameIdx := reader.readUnit16()
	attrName := cp.getUtf8(attrNameIdx)

	//读取attribute_length
	attrLen := reader.readUnit32()

	//读取u1[attribute_length]
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)

	return attrInfo
}

/**
1. 先读出了attributeName和attributeLength
2. 然后根据attributeName去生成分别的attributeInfo
也就是说每一个attribute的构成的前两个都是u2的name index和u4的长度
*/
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{
			cp: cp,
		}
	case "Exceptions":
		return &ExceptionAttributes{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{
			cp: cp,
		}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	default:
		return &UnParsedAttribute{
			name:   attrName,
			length: attrLen,
			info:   nil,
		}
	}
}
