package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func parseCmd() *Cmd {
	Cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&Cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&Cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&Cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&Cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&Cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		Cmd.class = args[0]
		Cmd.args = args[1:]
	}

	return Cmd
}
