package fileutils

import "os"


func WriteToFile(filename string, data string) error {
    return os.WriteFile(filename, []byte(data), 0644)
}
