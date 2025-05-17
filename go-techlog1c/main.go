package main

import (
    
    "techLog1C\database"
    
)

func main(){
    
    // считываем конфиг
    var config conf
    config.getConfig()
    
    // подключаем логи
    initLogging(&config)
    deleteOldLogFiles(&config)
    
    // maxdop установка
    runtime.GOMAXPROCS(config.MaxDop)
    
    // получаем файлы логов, сортируем по размеру, определяем в пакеты заданий
    arr,err: = getFilesArray(config.Path)
    if err != nil {
        logr.WithFields(logr.Fields {
            "object": "Data",
            "title": "Failure to scan directory",
        }).Error(err)
    }
    
}