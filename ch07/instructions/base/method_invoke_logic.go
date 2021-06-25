package base

import (
	"go-jvm/ch07/rtda"
	"go-jvm/ch07/rtda/heap"
)

//frame 中push进去method的栈帧
func InvokeMethod(frame *rtda.Frame, method *heap.Method)  {
	slotCount := method.ArgSlotCount()
	stack := frame.OperandStack()

	newFrame := frame.Thread().NewFrame(method)
	frame.Thread().PushFrame(newFrame)

	newLocalVars := newFrame.LocalVars()

	if slotCount > 0 {
		for i := slotCount - 1; i >= 0; i-- {
			newLocalVars.SetSlot(uint(i), stack.PopSlot())
		}
	}

}