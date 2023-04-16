package info

import (
	"fmt"

	"github.com/peterargue/flow-info/internal"
)

// Save downloads a file from a url and saves it to a file.
func Save(url, saveTo string) error {
	data, err := internal.Download(url)
	if err != nil {
		return fmt.Errorf("error downloading data: %w", err)
	}

	return internal.WriteFile(saveTo, data)
}
