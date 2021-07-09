package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
	"math"
)

func init() {
	//public static native long doubleToRawLongBits(double value);
	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)

	//public static native double longBitsToDouble(long bits);
	native.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

func longBitsToDouble(frame *rtda.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.OperandStack().PushDouble(value)
}

func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}
