package bufio

import (
	"bufio"
	"io"
	"readersandwriters/printer"
	"strings"
)

// Буферизированное чтение
type CustomReader struct {
	reader    io.Reader
	readCount int
}

func NewCustomReader(reader io.Reader) *CustomReader {
	return &CustomReader{reader, 0}
}

func (cr *CustomReader) Read(slice []byte) (count int, err error) {
	count, err = cr.reader.Read(slice)
	cr.readCount++
	printer.Printfln("Custom Reader: %v bytes", count)
	if err == io.EOF {
		printer.Printfln("Total Reads: %v", cr.readCount)
	}
	return
}

// Буферизированная запись
type CustomWriter struct {
	writer     io.Writer
	writeCount int
}

func NewCustomWriter(writer io.Writer) *CustomWriter {
	return &CustomWriter{writer, 0}
}
func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
	count, err = cw.writer.Write(slice)
	cw.writeCount++
	printer.Printfln("Custom Writer: %v bytes", count)
	return
}
func (cw *CustomWriter) Close() (err error) {
	if closer, ok := cw.writer.(io.Closer); ok {
		closer.Close()
	}
	printer.Printfln("Total Writes: %v", cw.writeCount)
	return
}

func Read() {

	printer.PrintTotal("Read Bufio:")

	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)

	buffered := bufio.NewReader(reader)

	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			printer.Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	printer.Printfln("Read data: %v", writer.String())

}

func Write() {

	text := "It was a boat. A small boat."
	printer.PrintTotal("Write Bufio:")

	var builder strings.Builder
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)

	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	printer.Printfln("Written data: %v", builder.String())

}
