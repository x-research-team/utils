package directory

import (
	"os"
	"path/filepath"
)

var Cwd = func() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	root, err := filepath.Abs(cwd)
	if err != nil {
		panic(err)
	}
	return root
}()
