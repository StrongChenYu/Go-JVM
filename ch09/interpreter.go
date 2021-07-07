package main

import (
	"fmt"
	"go-jvm/ch09/instructions"
	"go-jvm/ch09/instructions/base"
	"go-jvm/ch09/rtda"
	"go-jvm/ch09/rtda/heap"
)

func interpret(method *heap.Method, logInst bool, args []string) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	class := method.Class()

	jArgsRef := goArraysToJStr(class.Loader(), args)
	frame.LocalVars().SetRef(0, jArgsRef)

	defer catchErr(thread)
	loop(thread, logInst)
}

func goArraysToJStr(loader *heap.ClassLoader, goArgs []string) *heap.Object {
	class := loader.LoadClass("java/lang/String")
	jArgRef := class.ArrayClass().NewArray(uint(len(goArgs)))

	refs := jArgRef.Refs()
	for i, goArg := range goArgs {
		refs[i] = heap.JString(loader, goArg)
	}

	return jArgRef
}

//捕获异常
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.ByteCodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			//这里会打印操作数
			logInstruction(frame, inst)
		}

		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	//打印格式
	//InvokeDemo.main() #10 *control.RETURN &{{}}
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
