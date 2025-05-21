package main

import (
	"bufio"
	"fmt"
	"os"

	assembler "github.com/RISC-GoV/risc-assembler"
)

func runRepl() {
	reader := bufio.NewReader(os.Stdin)
	asm := assembler.Assembler{}
	for {
		fmt.Print("$>")
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			break
		}
		code, err := asm.AssembleLine(text)
		if err != nil {
			fmt.Println(err)
		}
		for b := range code {
			fmt.Printf("%08b", b)
		}
		fmt.Println()
	}
	fmt.Println("Goodbye.")
}
