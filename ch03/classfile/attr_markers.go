package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}


type MarkerAttribute struct {
}

func (m MarkerAttribute) readInfo(reader *ClassReader) {
	//只是标记的属性，什么也不做
}



