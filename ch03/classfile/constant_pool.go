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

	}

	return cp
}

func (receiver ConstantPool) getConstantInfo(index uint16) ConstantInfo  {
	if cpInfo := receiver[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (receiver ConstantPool) getNameAndType(index uint16) (string, string) {
	return "",""
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
	return nil
}