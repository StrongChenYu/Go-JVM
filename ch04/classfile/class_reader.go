package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (receiver *ClassReader) readUnit8() uint8 {
	val := receiver.data[0]
	receiver.data = receiver.data[1:]
	return val
}

func (receiver *ClassReader) readUnit16() uint16  {
	val := binary.BigEndian.Uint16(receiver.data)
	receiver.data = receiver.data[2:]
	return val
}

func (receiver *ClassReader) readUnit32() uint32  {
	val := binary.BigEndian.Uint32(receiver.data)
	receiver.data = receiver.data[4:]
	return val
}

func (receiver *ClassReader) readUnit64() uint64 {
	val := binary.BigEndian.Uint64(receiver.data)
	receiver.data = receiver.data[8:]
	return val
}

func (receiver *ClassReader) readUnit16s() []uint16 {
	n := receiver.readUnit16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = receiver.readUnit16()
	}
	return s
}

func (receiver *ClassReader) readBytes(length uint32) []byte {
	s := receiver.data[:length]
	receiver.data = receiver.data[length:]
	return s
}