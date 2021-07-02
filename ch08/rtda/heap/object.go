package heap

type Object struct {
	//用于存放引用类型
	class  *Class
	fields Slots
}

func (o Object) Class() *Class {
	return o.class
}

func (o Object) Fields() Slots {
	return o.fields
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(o.class)
}
