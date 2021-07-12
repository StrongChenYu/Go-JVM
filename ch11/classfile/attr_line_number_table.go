package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
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

//举例说明
//一行代码必定对应1-n个字节码
//而startpc就是字节码的起始位置
//那么
//1. originCode1 	bytecode  startPc
//                  bytecode
//                  bytecode
//                  bytecode
//                  bytecode
//2. originCode2 	bytecode  startPC
//                  bytecode
//                  bytecode
//                  bytecode
//                  bytecode
//                  bytecode
//                  bytecode
//这样一来就很明显了，倒着往回找，找到第一个大于等于pc的startPC
//为什么传入的是frame的nextpc-1
//因为这样的话起始找到的就是出错的那一行代码的lineNumber
func (self *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(self.lineNumberTable) - 1; i >= 0; i-- {
		entry := self.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}
