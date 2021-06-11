package classfile

/**
	本地变量表
 */
type LocalVariableTableAttribute struct {
	LocalVariableTable		[]*LocalVariableTableEntry
}

/**
	读取本地变量表
 */
func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUnit16()
	l.LocalVariableTable = make([]*LocalVariableTableEntry, length)

	for i := range l.LocalVariableTable {
		l.LocalVariableTable[i] = &LocalVariableTableEntry{
			startPc:       reader.readUnit16(),
			length:        reader.readUnit16(),
			nameIdx:       reader.readUnit16(),
			descriptorIdx: reader.readUnit16(),
			index:         reader.readUnit16(),
		}
	}
}


type LocalVariableTableEntry struct {
	startPc			uint16
	length			uint16
	nameIdx			uint16
	descriptorIdx	uint16
	index 			uint16
}
