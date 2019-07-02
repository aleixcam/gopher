package board

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// WriteOutput writes the file with the output
func WriteOutput(path string, messages []string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "Impossible open file")
	}

	var output []byte
	for _, message := range messages {
		output = append(output, fmt.Sprintf("<p>%s</p>", message)...)
	}

	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return errors.Wrapf(err, "Impossible open file")
	}

	defer f.Close()

	return nil
}
