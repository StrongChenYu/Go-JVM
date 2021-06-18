package classfile

type ConstantClassInfo struct {
	cp			ConstantPool
	nameIndex	uint16
}

func (ci *ConstantClassInfo) readInfo(reader *ClassReader) {
	ci.nameIndex = reader.readUnit16()
}

func (ci *ConstantClassInfo) Name() string {
	return ci.cp.getUtf8(ci.nameIndex)
}



