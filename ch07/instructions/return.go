package instructions

import (
	"go-jvm/ch07/instructions/base"
	"go-jvm/ch07/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame)  {

}
