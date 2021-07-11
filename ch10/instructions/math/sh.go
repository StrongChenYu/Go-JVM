package math

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/rtda"
)

//移位运算
type ISHL struct{ base.NoOperandsInstruction }

//这个是有符号算术右移
type ISHR struct{ base.NoOperandsInstruction }

//这个是无符号逻辑右移
type IUSHR struct{ base.NoOperandsInstruction }

type LSHL struct{ base.NoOperandsInstruction }

//这个是有符号算术右移
type LSHR struct{ base.NoOperandsInstruction }

//这个是无符号逻辑右移
type LUSHR struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	//只取后五位，因为五位表达的值足够了，剩下的位被忽略
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	//go 语言中没有>>>指令
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	//只取后五位，因为五位表达的值足够了，剩下的位被忽略
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	//go 语言中没有>>>指令
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
