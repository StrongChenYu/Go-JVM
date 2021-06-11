package classfile

type LineNumberTableAttribute struct {
	lineNumberTable 		[]*LineNumberTableEntry
}

//读取lineNumber Entry
func (l *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUnit16()
	l.lineNumberTable = make([]*LineNumberTableEntry, length)

	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUnit16(),
			lineNumber: reader.readUnit16(),
		}
	}
}

type LineNumberTableEntry struct {
	startPc			uint16
	lineNumber 		uint16
}

