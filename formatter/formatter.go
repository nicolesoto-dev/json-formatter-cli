package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrint(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "  ")
	return out.Bytes(), err
}

func Minify(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Compact(&out, data)
	return out.Bytes(), err
}

func hello() {
	fmt.Println("Hello! :)")
}
