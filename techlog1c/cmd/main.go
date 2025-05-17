package main

import (
	"techLog1C/techlog"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	logr "github.com/sirupsen/logrus"
)

func main() {

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yml")
	if err != nil {
		panic(err)
	}

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
