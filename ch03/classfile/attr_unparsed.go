package classfile

//没有解析的属性
type UnParsedAttribute struct {
	name 		string
	length		uint32
	info 		[]byte
}

func (u *UnParsedAttribute) readInfo(reader *ClassReader) {
	u.info = reader.readBytes(u.length)
}

