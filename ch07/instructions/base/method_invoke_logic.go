package base

import (
	"fmt"
	"go-jvm/ch07/rtda"
	"go-jvm/ch07/rtda/heap"
)

//frame 中push进去method的栈帧
func InvokeMethod(frame *rtda.Frame, method *heap.Method)  {
	slotCount := method.ArgSlotCount()
	stack := frame.OperandStack()
	thread := frame.Thread()

	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	newLocalVars := newFrame.LocalVars()

	if slotCount > 0 {
		for i := slotCount - 1; i >= 0; i-- {
			newLocalVars.SetSlot(uint(i), stack.PopSlot())
		}
	}

	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}