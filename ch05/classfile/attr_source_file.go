package classfile

//SourceFile的属性
type SourceFileAttribute struct {
	cp					ConstantPool
	sourceFileIdx		uint16
}

func (s *SourceFileAttribute) readInfo(reader *ClassReader) {
	s.sourceFileIdx = reader.readUnit16()
}

func (s *SourceFileAttribute) FileName() string {
	return s.cp.getUtf8(s.sourceFileIdx)
}

