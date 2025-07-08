package formatter

import (
	// "bytes"
	// "encoding/json"
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
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return "", fmt.Errorf("failed to indent JSON: %w", err)
	}
	return prettyJSON.String(), nil
}
