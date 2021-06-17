package main

import (
	"fmt"
	"go-jvm/ch05/classfile"
	"go-jvm/ch05/instructions"
	"go-jvm/ch05/instructions/base"
	"go-jvm/ch05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo)  {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, byteCode)
}


//捕获异常
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		inst.Execute(frame)

		fmt.Printf("pc %2d inst %T %v\n", pc, inst, inst)

	}
}


