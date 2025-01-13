package pipe

import (
	"io"
	"readersandwriters/printer"
)

func generateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4
	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		printer.Printfln("Wrote  %v  byte(s):  %v", count, string(data[i:end]))
		if err != nil {
			printer.Printfln("Error: %v", err.Error())
		}
	}
	if closer, ok := writer.(io.Closer); ok {
		closer.Close()
	}

}

func consumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			printer.Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
		if err == io.EOF {
			break
		}
	}
	printer.Printfln("Read data: %v", string(data))
}

func ReadWrite() {

	printer.PrintTotal("Read\\Write String:")
	pipeReader, pipeWriter := io.Pipe()
	go generateData(pipeWriter)
	consumeData(pipeReader)
}
