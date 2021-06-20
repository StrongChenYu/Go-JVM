package heap

import (
	"fmt"
	"go-jvm/ch06/classfile"
)

type Constant interface {}

//常量池
//注意是运行时常量池
type ConstantPool struct {
	class 		*Class
	consts		[]Constant
}

//根据class文件中的constant解析到虚拟机中的常量池中
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	
	rtCp := &ConstantPool{
		class:  class,
		consts: consts,
	}
	
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Val()
		case *classfile.ConstantFloatInfo:
			floadInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floadInfo.Val()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Val()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Val()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
			//
			//case *classfile.ConstantClassInfo:
			//classInfo := cpInfo.(*classfile.ConstantClassInfo)
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp, fieldInfo)
		case *classfile.ConstantMethodRefInfo:
			methodInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			interfaceInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, interfaceInfo)
		}
	}

	return rtCp
}


//????????????这里索引和类文件里面的常量池索引一样吗？？？？？？？
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}