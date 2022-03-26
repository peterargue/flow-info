package info

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Download(url string, outfile string) error {
	client := http.Client{
		Timeout: time.Second * 120,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error getting data: %w", err)
	}

	if res.Body == nil {
		return fmt.Errorf("error getting data: empty body")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data: %w", err)
	}

	file, err := os.Create(outfile)
	if err != nil {
		return fmt.Errorf("error creating data: %w", err)
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		return fmt.Errorf("error writing data: %w", err)
	}

	return nil
}
