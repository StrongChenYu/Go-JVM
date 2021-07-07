package references

import (
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.ByteCodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	classRef := frame.Method().Class().ConstantPool().GetConstant(uint(self.index)).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	counts := popAndCheckCounts(stack, int(self.dimensions))
	arrRef := newMultiDimensionalArray(counts, class)
	stack.PushRef(arrRef)
}

func popAndCheckCounts(stack *rtda.OperandStack, dimension int) []int32 {
	counts := make([]int32, dimension)
	for i := dimension - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

//逐层去加载数组
//然后需要加载对应的类
func newMultiDimensionalArray(counts []int32, class *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := class.NewArray(count)

	//arr已经是最底层了
	//newArray就会初始化到最后一层的数据
	if len(counts) > 1 {
		arrRef := arr.Refs()
		for i := range arrRef {
			arrRef[i] = newMultiDimensionalArray(counts[1:], class.ComponentClass())
		}
	}
	return arr
}
