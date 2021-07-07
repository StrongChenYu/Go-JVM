package references

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

//与put_static一样,给类的某个字段赋值
type PUT_FIELD struct{ base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()
	class := field.Class()

	//操作数异常
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		curMethod := frame.Method()
		curClass := curMethod.Class()

		//public class Test{
		//	private final int x;
		//	public Test(int x){
		//		this.x = x;
		//	}
		//}
		//condition1: class != curClass only can init final field belong to self class
		//condition2: curMethod.Name() can only be init
		if class != curClass || curMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlodId()

	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	}
}
