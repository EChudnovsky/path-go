package main

import (
	"readersandwriters/bufio"
	"readersandwriters/pipe"
	"readersandwriters/str"
)

func main() {

	// чтение\запись строки
	str.ReadWrite()

	// чтение\запись через поток
	pipe.ReadWrite()

	// чтение через буфер
	bufio.Read()

	// запись через буфер
	bufio.Write()

}
