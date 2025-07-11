package formatter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FlagOpts struct {
	Validate bool
	Minify   bool
	Prettify bool
	Output   bool
}

const baseDir = "~"

func ReadFile(filePath string, opts FlagOpts) (string, error) {
	cleanPath := filepath.Join(baseDir, filepath.Clean("/"+filePath))

	if !strings.HasPrefix(cleanPath, baseDir) {
		return "", errors.New("invalid file path")
	}

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
			return "", fmt.Errorf("invalid JSON: %w", err)
		}
	}

	if opts.Validate && !opts.Minify && !opts.Prettify {
		return "JSON is valid", nil
	}

	switch {
	case opts.Minify:
		err := json.Compact(&prettyJSON, data)
		if err != nil {
			return "", fmt.Errorf("failed to prettify JSON: %w", err)
		}
	case opts.Prettify:
		err := json.Indent(&prettyJSON, data, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to prettify JSON: %w", err)
		}
	}
	return prettyJSON.String(), nil
}
