package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/x-research-team/utils/directory"
)

func Read(dir, filename string, v interface{}) error {
	f := fmt.Sprintf("%s/%s", path.Join(directory.Cwd, dir), filename)
	buffer, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buffer, &v); err != nil {
		return err
	}
	return nil
}
