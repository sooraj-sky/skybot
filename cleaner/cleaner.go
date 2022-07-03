package cleaner

import (
	"os"
	"path/filepath"
)

func RemoveAudioFIles(path string) (err error) {
	contents, err := filepath.Glob(path)
	if err != nil {
		return
	}
	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return
		}
	}
	return
}
