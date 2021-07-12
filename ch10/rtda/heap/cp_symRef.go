package heap

//注意：ClassRef并没有给class指针赋值，解析的目的正在于此
//解析的目的是解析class
//目的是cp.class -> class
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

//解析类符号引用：
//1. 加载这个类
//2. 判断对这个类有没有访问权限
//3. 保存解析引用，防止重复解析
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	//d是常量池
	//self是class d的常量池符号引用
	//使用d的类加载器去加载c
	c := d.loader.LoadClass(self.className)
	if !c.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
