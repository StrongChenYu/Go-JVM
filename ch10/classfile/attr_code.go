package classfile

/**
描述属性表中的code属性
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (c *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return c.exceptionTable
}

func (c *CodeAttribute) Code() []byte {
	return c.code
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	//读取maxStack和maxLocals
	c.maxStack = reader.readUnit16()
	c.maxLocals = reader.readUnit16()

	//读取code长度
	codeLength := reader.readUnit32()
	c.code = reader.readBytes(codeLength)

	//读取ExceptionTable
	c.exceptionTable = readExceptionTable(reader)

	//读取code中的attribute表
	c.attributes = readAttributes(reader, c.cp)
}

func (c *CodeAttribute) MaxStack() uint16 {
	return c.maxStack
}

func (c *CodeAttribute) MaxLocals() uint16 {
	return c.maxLocals
}

/**
读取exception表
1.先读取长度
2.然后读取每一个exception表
*/
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionLength := reader.readUnit16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUnit16(),
			endPc:     reader.readUnit16(),
			handlerPc: reader.readUnit16(),
			catchType: reader.readUnit16(),
		}
	}

	return exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (e ExceptionTableEntry) StartPc() uint16 {
	return e.startPc
}

func (e ExceptionTableEntry) EndPc() uint16 {
	return e.endPc
}

func (e ExceptionTableEntry) CatchType() uint16 {
	return e.catchType
}

func (e ExceptionTableEntry) HandlerPc() uint16 {
	return e.handlerPc
}
