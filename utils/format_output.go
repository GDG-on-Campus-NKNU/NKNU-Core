package utils

import "encoding/base64"

func FormatBase64Output(data string, err error) string {
	dataOutput := `""`
	var errOutput string
	if data != "" {
		dataOutput = data
	}
	if err != nil {
		errOutput = err.Error()
	}
	output := `{"data":` + dataOutput + `,"err":"` + errOutput + `"}`
	return base64.StdEncoding.EncodeToString([]byte(output))
}
