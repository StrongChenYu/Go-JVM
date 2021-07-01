package classfile

import "fmt"

type ClassFile struct {
	magicNumber			uint32
	minorVersion		uint16
	majorVersion		uint16
	constantPool 		ConstantPool
	accessFlags			uint16
	thisClass			uint16
	superClass			uint16
	interfaces			[]uint16
	fields				[]*MemberInfo
	methods				[]*MemberInfo
	attributes			[]AttributeInfo
}

//将二进制文件解析为ClassFile类
func Parse(classData []byte) (cf *ClassFile, err error)  {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//因为magic和version不需要验证
func (cf *ClassFile) read(reader *ClassReader)  {
	//读魔数
	cf.readAndCheckMagic(reader)
	//读版本号minor和major
	cf.readAncCheckVersion(reader)
	//读取常量池
	cf.constantPool = readConstantPool(reader)
	//读取访问标志
	cf.accessFlags = reader.readUnit16()
	//读取类
	cf.thisClass = reader.readUnit16()
	//读取superClass
	cf.superClass = reader.readUnit16()
	//读取interfaces
	cf.interfaces = reader.readUnit16s()
	//读取field
	cf.fields = readMembers(reader, cf.constantPool)
	//读取method
	cf.methods = readMembers(reader, cf.constantPool)
	//读取attribute
	cf.attributes = readAttributes(reader, cf.constantPool)
}

//check
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader)  {
	magic := reader.readUnit32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAncCheckVersion(reader *ClassReader)  {
	cf.minorVersion = reader.readUnit16()
	cf.majorVersion = reader.readUnit16()
	switch cf.majorVersion {
		case 45:
			return
		case 46,47,48,49,50,51,52:
			if cf.minorVersion == 0 {
				return
			}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}


func (cf *ClassFile) Attributes() []AttributeInfo {
	return cf.attributes
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Interfaces() []uint16 {
	return cf.interfaces
}

func (cf *ClassFile) SuperClass() uint16 {
	return cf.superClass
}

func (cf *ClassFile) ThisClass() uint16 {
	return cf.thisClass
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	names := make([]string, len(cf.interfaces))
	for i := range names {
		names[i] = cf.constantPool.getClassName(cf.interfaces[i])
	}
	return names
}