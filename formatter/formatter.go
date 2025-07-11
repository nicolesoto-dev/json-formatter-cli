package formatter

import (
	"bytes"
	"encoding/json"
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

func ReadFile(filePath string, opts FlagOpts) (string, error) {
	baseDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get working directory: %w", err)
	}

	absoluteFilePath, err := filepath.Abs(filepath.Clean(filePath))
	if err != nil {
		return "", fmt.Errorf("invalid file path: %w", err)
	}

	if !strings.HasPrefix(absoluteFilePath, baseDir) {
		return "", fmt.Errorf("access to path %q is not allowed", absoluteFilePath)
	}

	data, err := os.ReadFile(absoluteFilePath)
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
