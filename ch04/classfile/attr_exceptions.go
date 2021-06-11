package classfile

type ExceptionAttributes struct {
	exceptionIdxTable		[]uint16
}

/**
	读取方法中的异常
	1.ExceptionEntry中存储的是code中的异常
 */
func (e *ExceptionAttributes) readInfo(reader *ClassReader) {
	e.exceptionIdxTable = reader.readUnit16s()
}

func (e *ExceptionAttributes) ExceptionIdxTable() []uint16 {
	return e.exceptionIdxTable
}

