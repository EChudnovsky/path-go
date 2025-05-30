package main

import (
    "os"
    "io/ioutil"
    "gopkg.in/yaml.v3"
    logr "github.com/sirupsen/logrus"
    "log"
)

type conf struct {
    Path string `yaml:"path"`
    TechLogDetailsEvents string `yaml:"tech_log_details_events"`
    MaxDop int `yaml:"maxdop"`
    Sorting int `yaml:"sorting"`
    PathLogFile string `yaml:"path_logfile"`
    LogLevel int `yaml:"log_level"`
    LogLifeSpan int `yaml:"log_life_span"`
    DeleteTabsInContexts bool `yaml:"delete_tabs_in_contexts"`
    DeletePostfixInNameVirtualTables bool `yaml:"delete_postfix_in_name_virtual_tables"`
    InsecureSkipVerify bool `yaml:"skip_verify_certificates"`
}

func (c *conf) getConfig() *conf {

    yamlFile,
    err: = ioutil.ReadFile("./conf/settings.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func initLogging(c *conf) {

    year,
    month,
    day: = time.Now().Date()
    if _,
    err: = os.Stat(c.PathLogFile); os.IsNotExist(err) {
        os.Mkdir(c.PathLogFile, 2)
    }

    logFileName: = c.PathLogFile + "techLog1C_" + strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day) + ".json"
    fileLog,
    err: = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

    if err != nil {
        log.Fatal(err)
    }

    logr.SetFormatter(&logr.JSONFormatter {})
    logr.SetOutput(fileLog)

}

func deleteOldLogFiles(c *conf) {

    lifeSpan: = c.LogLifeSpan
    if lifeSpan == 0 {
        lifeSpan = 1
    }

    files,
    err: = getFilesArray(c.PathLogFile)
    if err != nil {
        logr.WithFields(logr.Fields {
            "object": "Scan log directory",
            "title": "Select life span files",
        }).Warning(err)
        return
    }

    t: = time.Now()

    for _,
    file: = range files {
        fname: = filepath.Base(file.Path)
        if !strings.HasPrefix(fname, "techLog1C_") {
            // избегает удаления не наших лог файлов, например при неверном указании пользователем нашего лог каталога
            continue
        }
        duration: = t.Sub(file.DataCreate)

        if (int(duration.Hours()) - 24*c.LogLifeSpan) >= 0 {
            if err = os.Remove(file.Path); err != nil {
                logr.WithFields(logr.Fields {
                    "object": "Scan log directory",
                    "title": "Delete life span files",
                }).Warning(err)
            }
        }
    }
}
