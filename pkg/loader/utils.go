package loader

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func getComposeFileDir(inputFiles []string) (string, error) {
	inputFile := inputFiles[0]
	if strings.Index(inputFile, "/") != 0 {
		workDir, err := os.Getwd()
		if err != nil {
			return "", errors.Wrap(err, "Unable to retrieve compose file directory")
		}
		inputFile = filepath.Join(workDir, inputFile)
	}

	return filepath.Dir(inputFile), nil
}
