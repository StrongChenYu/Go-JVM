package references

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
	"go-jvm/ch08/rtda/heap"
)

//给一个static变量赋值
type PUT_STATIC struct{ base.Index16Instruction }

//public class X {
//	public void test() {
//		Y.field = "chenyu"
//	}
//}
// X的constantPool中存储着Y.field的符号引用
// 所以第一步先解析了Y.field的符号引用
// Y.field的值存储在Y的class文件里面
func (self *PUT_STATIC) Execute(frame *rtda.Frame) {

	fieldRef := frame.Method().Class().ConstantPool().GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolveField()
	class := field.Class()

	//类初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	//操作数异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		curMethod := frame.Method()
		curClass := curMethod.Class()

		//condition1: class != curClass only can init final field belong to self class
		//condition2: curMethod.Name() can only be clinit
		if class != curClass || curMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlodId()
	staticVars := class.StaticVars()

	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		staticVars.SetInt(slotId, stack.PopInt())
	case 'J':
		staticVars.SetLong(slotId, stack.PopLong())
	case 'F':
		staticVars.SetFloat(slotId, stack.PopFloat())
	case 'D':
		staticVars.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		staticVars.SetRef(slotId, stack.PopRef())
	}
}
