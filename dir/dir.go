package dir

import (
	"os"

	"golang.org/x/exp/slices"
)

// CleanDirsExcept takes all folders that are childen of "path" and removes them
// recursively, without removing any dir listed in "exceptions"
func CleanDirsExcept(path string, exceptions []string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// already clean path
	if len(files) == 0 {
		return nil
	}

	// remove all folders except the "exception" folder
	for _, file := range files {
		if file.IsDir() && !slices.Contains(exceptions, file.Name()) {
			err := os.RemoveAll(path + "/" + file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
