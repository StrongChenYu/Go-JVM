package lang

import (
	"go-jvm/ch10/native"
	"go-jvm/ch10/rtda"
	"go-jvm/ch10/rtda/heap"
)

func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)

}

func arraycopy(frame *rtda.Frame) {
	localVars := frame.LocalVars()

	src := localVars.GetRef(0)
	srcPos := localVars.GetInt(1)
	dest := localVars.GetRef(2)
	destPos := localVars.GetInt(3)
	length := localVars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	//check array type
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	//check index
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, srcPos, dest, destPos, length)

}

func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}

	if srcClass.ComponentClass().IsPrimitive() && destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}

	return true
}
