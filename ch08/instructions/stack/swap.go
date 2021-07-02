package stack

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/rtda"
)

//交换位于栈顶的两个值的位置
//....[a][b]->
//....[b][a]->
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
