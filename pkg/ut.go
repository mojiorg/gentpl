package pkg

import (
	"os"
	"path/filepath"
)

func Chk(err error) {
	if err != nil {
		panic(err)
	}
}

func SaveFile(dirPath, fileName string, text []byte) error {
	file, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	/*p, err := format.Source(text)
	if err != nil {
		return err
	}
	fmt.Println(p)*/
	_, err = file.Write(text)
	return err
}

func MkdirPathIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, os.ModePerm)
	}
	return nil
}

func CleanUpGenFiles(dir string) error {
	exist, err := FileExists(dir)
	if err != nil {
		return err
	}
	if exist {
		return os.RemoveAll(dir)
	}
	return nil
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}
