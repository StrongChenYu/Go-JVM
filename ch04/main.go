package main

import (
	"fmt"
	"go-jvm/ch03/classfile"
	"go-jvm/ch03/classpath"
	"go-jvm/ch04/rtda"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100,100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testOperandStack(stack *rtda.OperandStack) {
	stack.PushInt(100)
	stack.PushInt(-100)
	stack.PushLong(2997924580)
	stack.PushLong(-2997924580)
	stack.PushFloat(3.1415926)
	stack.PushDouble(2.71828182845)
	stack.PushRef(nil)

	fmt.Println(stack.PopRef())
	println(stack.PopDouble())
	println(stack.PopFloat())
	println(stack.PopLong())
	println(stack.PopLong())
	println(stack.PopInt())
	println(stack.PopInt())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)

	fmt.Println(vars.GetInt(0))
	fmt.Println(vars.GetInt(1))
	fmt.Println(vars.GetLong(2))
	fmt.Println(vars.GetLong(4))
	fmt.Println(vars.GetFloat(6))
	fmt.Println(vars.GetDouble(7))
	fmt.Println(vars.GetRef(9))
}


func loadClass(class string, cp *classpath.Classpath) *classfile.ClassFile {
	//读取类数据
	classData, _, err := cp.ReadClass(class)
	if err != nil {
		panic(err)
	}

	//解析class数据
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}


//打印类相关的信息
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceName())

	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("    %s\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}

}