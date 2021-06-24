package main

import (
	"fmt"
	"go-jvm/ch07/classpath"
	"go-jvm/ch07/rtda/heap"
	"strings"
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

//?????????????加载别的类会报错？？？记得解决一下
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	class := classLoader.LoadClass(className)

	mainMethod := class.GetMainMethod()

	interpret(mainMethod)
}


