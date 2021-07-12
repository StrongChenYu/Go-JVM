package lang

import (
	"go-jvm/ch10/native"
	"go-jvm/ch10/rtda"
	"go-jvm/ch10/rtda/heap"
)

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func init() {
	//private native Throwable fillInStackTrace(int dummy);
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))

	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	return nil
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}