package classpath

type ClassReader struct {
	data []byte
}

func (receiver ClassReader) readUnit8() uint8 {
	return 0
}

func (receiver ClassReader) readUnit16() uint16  {
	return 0
}

func (receiver ClassReader) readUnit32() uint32  {
	return 0
}

func (receiver ClassReader) readUnit16s() []uint16 {
	return nil
}

func (receiver ClassReader) readBytes(length uint32) []byte {
	return nil
}