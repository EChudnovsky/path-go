package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func PrintTotal(template string) {
	Printfln("\n╔════════════════════════════════╗")
	Printfln("║%s 	", template)
	Printfln("╚════════════════════════════════╝")

}
