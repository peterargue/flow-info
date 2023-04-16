package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const downloadTimeout = time.Second * 120

func Download(url string) ([]byte, error) {
	client := http.Client{
		Timeout: downloadTimeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting data (url=%s): %w", url, err)
	}

	if res.Body == nil {
		return nil, fmt.Errorf("error getting data: empty body")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading data: %w", err)
	}

	return body, nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file (path=%s): %w", path, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file (path=%s): %w", path, err)
	}

	return data, nil
}

func WriteFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file (path=%s): %w", path, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing file (path=%s): %w", path, err)
	}

	return nil
}
