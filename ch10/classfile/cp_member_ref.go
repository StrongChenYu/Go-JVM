package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUnit16()
	c.nameAndTypeIndex = reader.readUnit16()
}

func (c *ConstantMemberRefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

//类中的字段
//private int a = 1;
//private int b = 2;
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

//普通方法
//public int printInfo()
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

//接口方法
//public abstract void printInfo()
//public void printInfo()
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}
