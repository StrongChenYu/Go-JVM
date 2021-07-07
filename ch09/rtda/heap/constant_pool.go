package heap

import (
	"fmt"
	"go-jvm/ch09/classfile"
)

type Constant interface{}

//常量池
//注意是运行时常量池
type ConstantPool struct {
	class  *Class
	consts []Constant
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
		//整数常量值
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Val()

			//浮点值
		case *classfile.ConstantFloatInfo:
			floadInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floadInfo.Val()

			//长整型值
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Val()
			i++
			//double
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Val()
			i++
			//字符串值
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
			//
			//case *classfile.ConstantClassInfo:
			//classInfo := cpInfo.(*classfile.ConstantClassInfo)

			//class符号引用
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
			//field符号引用
		case *classfile.ConstantFieldRefInfo:
			fieldInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp, fieldInfo)
			//method符号引用
		case *classfile.ConstantMethodRefInfo:
			methodInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodInfo)
			//method interface 符号引用
		case *classfile.ConstantInterfaceMethodRefInfo:
			interfaceInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, interfaceInfo)
		default:
			//todo
		}
	}

	return rtCp
}

//????????????这里索引和类文件里面的常量池索引一样吗？？？？？？？
//可以看到上面的default语句中什么都不做，而且索引也是从1开始的
//所以可以确定，classfile中的constantpool索引和
//这里的索引一样，如果是除了符号引用和4种常量值之外的值，对应的slot位置为空，不做处理
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
