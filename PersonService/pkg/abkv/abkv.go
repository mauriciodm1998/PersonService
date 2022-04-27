package abkv

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MurilloVaz/bitcask"
)

var (
	connections = map[string]*bitcask.Bitcask{}
)

const (
	pathString   = "%s/%s/"
	dbFileString = "%s%s"
)

func Open(name string, path string) (*bitcask.Bitcask, error) {
	if conn, ok := connections[name]; ok {
		return conn, nil
	}
	fullPath := filepath.Join(name, path)

	if alreadyExists(fullPath) {
		files, err := ioutil.ReadDir(fullPath)

		if err != nil {
			return nil, err
		}

		for _, f := range files {
			if fileName := f.Name(); strings.HasPrefix(fileName, "lock") {
				if err := os.Remove(fullPath + "/" + fileName); err != nil {
					return nil, err
				}
			}
		}
	}

	db, err := bitcask.Open(fullPath)

	if err != nil {
		return nil, err
	}

	connections[name] = db

	return db, nil
}

func alreadyExists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
