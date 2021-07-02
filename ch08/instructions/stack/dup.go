package stack

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

//????????这里会不会出问题，两个引用指向一个数值 所有的复制操作都是浅复制

//复制位于栈顶的一个值
type DUP struct{ base.NoOperandsInstruction }

//Duplicate the top operand stack value and insert two values down
//...[a][b][c]->
//...[a][c][b][c]->
type DUP_X1 struct{ base.NoOperandsInstruction }

//Duplicate the top operand stack value and insert two or three values down ?????? why not handler two situation
//...[a][b][c][d]->
//...[a][d][b][c][d]->
type DUP_X2 struct{ base.NoOperandsInstruction }

//复制位于栈顶的两个值
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (self *DUP) Execute(frame *rtda.Frame) {
	slot := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot)
	frame.OperandStack().PushSlot(slot)
}

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slotTop := stack.PopSlot()
	slotDownTop := stack.PopSlot()

	stack.PushSlot(slotTop)
	stack.PushSlot(slotDownTop)
	stack.PushSlot(slotTop)
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
