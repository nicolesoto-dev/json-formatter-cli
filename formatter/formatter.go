package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type FlagOpts struct {
	Validate bool
	Minify   bool
	Prettify bool
	Output   bool
}

func ReadFile(filePath string, opts FlagOpts) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return ProcessJSON(data, opts)
}

func ProcessJSON(data []byte, opts FlagOpts) (string, error) {
	var prettyJSON bytes.Buffer
	var v interface{}

	if opts.Validate {
		err := json.Unmarshal(data, &v)
		if err != nil {
			return "error:", err
		}
	}

	if opts.Validate && !opts.Minify && !opts.Prettify {
		return "JSON is valid", nil
	}

	if opts.Prettify {
		json.Indent(&prettyJSON, data, "", "  ")
	}

	if opts.Minify {
		json.Compact(&prettyJSON, data)
	}
	return prettyJSON.String(), nil
}
