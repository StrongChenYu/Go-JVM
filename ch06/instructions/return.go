package instructions

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame)  {

}
