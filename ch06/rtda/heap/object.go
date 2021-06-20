package heap

type Object struct {
	//用于存放引用类型
	class 		*Class
	fields 		Slots
}
