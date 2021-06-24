package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	//先读出常量池的数量
	cpCount := int(reader.readUnit16())
	//然后创建常量池
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		//这里是读出来
		//记得long和double要占用两个位置,所以记得跳过
		switch cp[i].(type) {
			case *ConstantLongInfo, *ConstantDoubleInfo:
				i++
		}
	}

	return cp
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}


func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo  {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

//获取字段的类型和名字
//int[] a = new int[10]
//name: a
//type: [I
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndTypeInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(nameAndTypeInfo.nameIndex)
	_type := cp.getUtf8(nameAndTypeInfo.typeIndex)
	return name, _type
}

//获取类的名字
//constantInfo[i]
func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUnit8()
	ci := newConstantInfo(tag, cp)
	ci.readInfo(reader)
	return ci
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{
			cp: cp,
		}
	case CONSTANT_FieldRef:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{
			cp: cp,
		}}
	case CONSTANT_MethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{
			cp: cp,
		}}
	case CONSTANT_InterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{
			cp: cp,
		}}
	case CONSTANT_String:
		return &ConstantStringInfo{
			cp: cp,
		}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	case CONSTANT_MethodHandler:
		return &ConstantMethodHandlerInfo{}
	//.......

	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}

}