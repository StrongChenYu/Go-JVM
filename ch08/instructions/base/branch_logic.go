package base

import "go-jvm/ch08/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
