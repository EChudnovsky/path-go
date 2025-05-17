package techlog

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

type record struct {
	event string
}

type files struct {
	Path          string
	Size          int64
	LastPosition  int64
	FileDate      string
	DataCreate    time.Time
	ProcessNameID string
	BlokingID     string
}

func ReadFiles((root string){

	// получаем файлы логов, сортируем по размеру, определяем в пакеты заданий
	arrFiles, err := techlog.GetFilesArray(config.String("Path"))
	if err != nil {
		logr.WithFields(logr.Fields{
			"object": "Data",
			"title":  "Failure to scan directory",
		}).Error(err)
	}

	for i := 0; i < len(arrFiles); i++ {
		// получаем дату из имени файла и тип процесса
		fileSplitter := strings.Split(arr[i].Path, separator)
		lenArray := len(fileSplitter)
		if lenArray < 2 {
			continue
		}

	}

}

func GetFilesArray(root string) ([]files, error) {

	var arrFiles []files

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			arrFiles = append(arrFiles, files{
				Path: path, Size: info.Size(), DataCreate: info.ModTime()})
		}
		return nil
	})
	return arrFiles, err
}

func readFile(file files) ([]byte, int64, error) {

	currPosition := file.LastPosition

	openFile, err := os.Open(file.Path)
	if err != nil {
		return nil, 0, err
	}

	buffer := make([]byte, 64)
	var data []byte

	openFile.Seek(file.LastPosition, 0)

	for {
		n, err := openFile.Read(buffer)
		currPosition += int64(n)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data = append(data, buffer[:n]...)
	}

	openFile.Close()
	return data, currPosition, nil
}
