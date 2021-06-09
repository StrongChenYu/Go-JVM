package classfile

type UnParsedAttribute struct {
	name 		string
	length		uint32
	info 		[]byte
}

func (u *UnParsedAttribute) readInfo(reader *ClassReader) {
	u.info = reader.readBytes(u.length)
}

