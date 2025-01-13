package str

import (
	"io"
	"readersandwriters/printer"
	"strings"
)

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			printer.Printfln("Read  %v  bytes:  %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
	printer.Printfln("String builder contents: %s", writer.(*strings.Builder).String())
}

func ReadWrite() {

	printer.PrintTotal("Read\\Write String")

	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)

}
