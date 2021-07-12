package heap

import "go-jvm/ch10/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

func newExceptionTable(exceptionEntries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	exceptionTable := make([]*ExceptionHandler, len(exceptionEntries))
	for i, entry := range exceptionEntries {
		exceptionTable[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(cp, entry.CatchType()),
		}
	}
	return exceptionTable
}

func getCatchType(cp *ConstantPool, index uint16) *ClassRef {
	if index == 0 {
		return nil
	}
	return cp.GetConstant(uint(index)).(*ClassRef)
}

func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}

			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
