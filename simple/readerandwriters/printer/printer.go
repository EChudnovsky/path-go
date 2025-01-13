package printer

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func PrintTotal(template string) {
	Printfln("_________________________")
	Printfln("	%s 	", template)
	Printfln("_________________________")

}
