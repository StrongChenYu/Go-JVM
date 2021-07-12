package classfile

//这部分是markers属性，不占用任何位置

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
