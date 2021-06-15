package stack

import (
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

type DUP struct { base.NoOperandsInstruction }
//Duplicate the top operand stack value and insert two values down
type DUP_X1 struct { base.NoOperandsInstruction }
type DUP_X2 struct { base.NoOperandsInstruction }
type DUP2 struct { base.NoOperandsInstruction }
type DUP2_X1 struct { base.NoOperandsInstruction }
type DUP2_X2 struct { base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtda.Frame)  {
	slot := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot)
	//???这里会不会出问题，两个引用指向一个数值
	frame.OperandStack().PushSlot(slot)
}

func (self *DUP_X1) Execute(frame *rtda.Frame)  {
	slotTop := frame.OperandStack().PopSlot()
	slotDownTop := frame.OperandStack().PopSlot();

	frame.OperandStack().PushSlot(slotTop)
	frame.OperandStack().PushSlot(slotDownTop)
	frame.OperandStack().PushSlot(slotTop)
}

func (self *DUP_X2) Execute(frame *rtda.Frame)  {

}
