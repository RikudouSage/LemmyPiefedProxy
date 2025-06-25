package helper

import (
	"encoding/json"
	"io"
	"strings"
)

func MarshalToReader(input any) io.Reader {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil
	}

	return strings.NewReader(string(bytes))
}
