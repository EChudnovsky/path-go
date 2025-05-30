package main

type record struct{
    event string
    
}

func getFilesArray(root string) ([]files, error) {
    var arrFiles []files

    err: = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            arrFiles = append(arrFiles, files {
                Path: path, Size: info.Size(), DataCreate: info.ModTime()})
        }
        return nil
    })
    return arrFiles, err
}

