package heap

import "go-jvm/ch08/classfile"

type Field struct {
	ClassMember
	constsValueIndex uint
	slotId           uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
		fields[i].copyAttributes(cfFields)
	}
	return fields
}

func (self *Field) IsStatic() bool {
	return self.accessFlags&ACC_STATIC != 0
}

func (self *Field) IsFinal() bool {
	return self.accessFlags&ACC_FINAL != 0
}

func (self *Field) IsLongOrDouble() bool {
	return self.descriptor == "D" || self.descriptor == "J"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constsValueIndex = uint(valAttr.ConstantValueIdx())
	}
}

func (self *Field) ConstsValueIndex() uint {
	return self.constsValueIndex
}

func (self *Field) SlodId() uint {
	return self.slotId
}
