package files

import (
	"io"
	"os"
)

func Exists(path string) bool {
	path = os.ExpandEnv(path)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Read(path string) ([]byte, error) {
	path = os.ExpandEnv(path)
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}

func Open(path string) (*os.File, error) {
	path = os.ExpandEnv(path)
	return os.Open(path)
}

func Remove(path string) error {
	path = os.ExpandEnv(path)
	return os.Remove(path)
}

func Write(path string, content []byte) error {
	path = os.ExpandEnv(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func Mkdir(path string) error {
	path = os.ExpandEnv(path)
	return os.MkdirAll(path, 0755)
}
