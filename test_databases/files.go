package testdatabases

import (
	"path"
	"runtime"
)

const (
	jet4DatabasePath = "EPSG_v8_6.mdb"
	jet3DatabasePath = "Books_be.mdb"
)

func dbPath(file string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), file)
}

func JET4DatabasePath() string {
	return dbPath(jet4DatabasePath)
}

func JET3DatabasePath() string {
	return dbPath(jet3DatabasePath)
}
