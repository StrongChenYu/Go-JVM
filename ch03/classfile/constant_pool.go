package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	return nil
}

func (receiver ConstantPool) getConstantInfo(index uint16) ConstantInfo  {
	return nil	
}

func (receiver ConstantPool) getNameAndType(index uint16) (string, string) {
	return "",""
}


type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	return nil
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	return nil
}