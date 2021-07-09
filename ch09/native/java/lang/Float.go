package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/rtda"
	"math"
)

func init() {
	//public static native int floatToRawIntBits(float value);
	native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
	//public static native float intBitsToFloat(int bits);
	native.Register("java/lang/Float", "intBitsToFloat", "(I)F", intBitsToFloat)
}

func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits))
	frame.OperandStack().PushFloat(value)
}

func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}
