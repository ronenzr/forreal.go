package FileUtil

import (
	"io/ioutil"
	"os"
)

func GetFileContent(path string) string {
	f, err := os.Open(path) // can't use a wrapper function or pkger doesn't pack it!
	if err != nil {
		panic(err) // we know that this file exists, so this is just to be sure and to get error messages that make sense
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err) // this will hopefully only possibly panic during development as the file is already in memory otherwise
	}
	return string(b)
}
