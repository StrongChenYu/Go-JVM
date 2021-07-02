package extended

import (
	"go-jvm/ch08/instructions/base"
	"go-jvm/ch08/instructions/loads"
	"go-jvm/ch08/instructions/math"
	"go-jvm/ch08/instructions/stores"
	"go-jvm/ch08/rtda"
)

//扩展指令
//因为localVars的大小最大默认是一个字节的大小256
//但是也有可能出现不够用的情况，所以这种情况下就需要对本地变量表的索引进行扩展
type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.ByteCodeReader) {
	opCode := reader.ReadUint8()
	switch opCode {
	case 0x15: //iload
		inst := &loads.ILOAD{}
		//索引所以是无符号数
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16: //load
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17: //fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18: //dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19: //aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36: //istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37: //lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38: //fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39: //dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a: // astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84: // iinc
		inst := &math.IINC{
			Index: uint(reader.ReadUint16()),
			Const: int32(reader.ReadInt16()),
		}
		self.modifiedInstruction = inst
	case 0xa9: //ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
